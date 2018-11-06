package txpool

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/iost-official/go-iost/core/tx"
	"github.com/iost-official/go-iost/db/kv"

	"github.com/emirpasic/gods/trees/redblacktree"
	"github.com/uber-go/atomic"
)

func compareDelayTx(a, b interface{}) int {
	txa := a.(*tx.Tx)
	txb := b.(*tx.Tx)
	if txa.Time == txb.Time {
		return bytes.Compare(txa.Hash(), txb.Hash())
	}
	return int(txa.Time - txb.Time)
}

// DeferServer manages delayed transaction and sends them to txpool on time.
type DeferServer struct {
	deferTxDB *kv.Storage

	index            *redblacktree.Tree
	rw               *sync.RWMutex
	nextScheduleTime atomic.Int64

	txpool *TxPImpl

	quitCh chan struct{}
}

// NewDeferServer returns a new DeferServer instance.
func NewDeferServer(path string, txpool *TxPImpl) (*DeferServer, error) {
	levelDB, err := kv.NewStorage(path, kv.LevelDBStorage)
	if err != nil {
		return nil, fmt.Errorf("fail to init deferserver DB, %v", err)
	}

	deferServer := &DeferServer{
		deferTxDB: levelDB,
		index:     redblacktree.NewWith(compareDelayTx),
		rw:        new(sync.RWMutex),
		txpool:    txpool,
		quitCh:    make(chan struct{}),
	}
	err = deferServer.buildIndex()
	if err != nil {
		return nil, fmt.Errorf("build defertx index error, %v", err)
	}

	return deferServer, nil
}

func (d *DeferServer) buildIndex() error {
	keys, err := d.deferTxDB.Keys([]byte{})
	if err != nil {
		return err
	}
	for _, key := range keys {
		newTx, err := d.getNewTx(key)
		if err != nil {
			return err
		}
		d.index.Put(newTx, true)
	}
	return nil
}

func (d *DeferServer) getNewTx(txHash []byte) (*tx.Tx, error) {
	newTxBytes, err := d.deferTxDB.Get(txHash)
	if err != nil {
		return nil, err
	}
	newTx := &tx.Tx{}
	err = newTx.Decode(newTxBytes)
	return newTx, err
}

// DelDelaytx deletes a tx in defer server.
func (d *DeferServer) DelDelaytx(txHash []byte) error {
	newTx, err := d.getNewTx(txHash)
	if err != nil {
		return err
	}
	d.deferTxDB.Delete(txHash)
	d.rw.Lock()
	d.index.Remove(newTx)
	d.rw.Unlock()
	return nil
}

// StoreDelaytx stores a tx in defer server.
func (d *DeferServer) StoreDelaytx(t *tx.Tx) {
	newTx := &tx.Tx{
		ReferredTx: t.ReferredTx,
		Time:       t.Time + t.Delay,
	}
	d.deferTxDB.Put(t.Hash(), newTx.Encode())
	d.rw.Lock()
	d.index.Put(newTx, true)
	d.rw.Unlock()
	if newTx.Time < d.nextScheduleTime.Load() {
		d.nextScheduleTime.Store(newTx.Time)
		d.restartDeferTicker()
	}
}

// DumpDelayTx dumps all delay transactions for debug.
func (d *DeferServer) DumpDelayTx() []*tx.Tx {
	ret := make([]*tx.Tx, 0)
	d.rw.RLock()
	d.rw.RUnlock()
	return ret
}

// Start starts the defer server.
func (d *DeferServer) Start() error {
	go d.deferTicker()
	return nil
}

// Stop stops the defer server.
func (d *DeferServer) Stop() {
	d.stopDeferTicker()
}

func (d *DeferServer) stopDeferTicker() {
	d.quitCh <- struct{}{}
	<-d.quitCh
}

func (d *DeferServer) restartDeferTicker() {
	d.stopDeferTicker()
	go d.deferTicker()
}

func (d *DeferServer) deferTicker() {
	for {
		select {
		case <-d.quitCh:
			d.quitCh <- struct{}{}
			return
		case <-time.After(time.Duration(d.nextScheduleTime.Load() - time.Now().UnixNano())):
			iter := d.index.Iterator()
			d.rw.RLock()
			ok := iter.Next()
			d.rw.RUnlock()
			for ok {
				newTx := iter.Key().(*tx.Tx)
				if newTx.Time > time.Now().UnixNano() {
					d.nextScheduleTime.Store(newTx.Time)
					break
				}
				err := d.txpool.AddDefertx(newTx)
				if err == nil || strings.Index(err.Error(), "DupError.") > 0 {
					d.rw.Lock()
					d.index.Remove(newTx)
					d.rw.Unlock()
				}
				d.rw.RLock()
				ok = iter.Next()
				d.rw.RUnlock()
			}
			if !ok {
				d.nextScheduleTime.Store(math.MaxInt64)
			}
		}
	}

}