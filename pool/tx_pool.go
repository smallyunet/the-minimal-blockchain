package pool

import (
	"github.com/smallyunet/tmb/block"
	"log"
	"time"

	"github.com/smallyunet/tmb/consensus"
)

var TxPool = make([]block.KeyValue, 0)

var DataMsg = make(chan int, 1)

func AcceptTx() {
	ticker := time.NewTicker(time.Duration(blockTime) * time.Millisecond)

	for {
		select {
		case <-DataMsg:
			if uint64(len(TxPool)) >= blockSize {
				pushBlockByTx()
			}
		case <-ticker.C:
			if len(TxPool) > 0 {
				pushBlockByTx()
			}
		default:
			// do nothing
		}
	}
}

func PushTxToPool(data block.KeyValue) {
	TxPool = append(TxPool, block.KeyValue{
		Key:   data.Key,
		Value: data.Value,
	})
}

func pushBlockByTx() {
	err := consensus.SaveBlockByTx(TxPool)
	if err != nil {
		log.Fatalln(err)
		return
	}
	TxPool = make([]block.KeyValue, 0)
}
