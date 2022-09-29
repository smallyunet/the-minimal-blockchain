package pool

import (
	"github.com/smallyunet/tmb/block"
	"log"
	"time"

	"github.com/smallyunet/tmb/consensus"
)

var PoolCache = make([]block.KeyValue, 0)

var DataMsg = make(chan int, 1)

func init() {
	go Accept()
}

func Accept() {
	ticker := time.NewTicker(time.Duration(blockTime) * time.Millisecond)

	for {
		select {
		case <-DataMsg:
			if uint64(len(PoolCache)) >= blockSize {
				pushBlock()
			}
		case <-ticker.C:
			if len(PoolCache) > 0 {
				pushBlock()
			}
		default:
			// do nothing
		}
	}
}

func PushToPool(data block.KeyValue) {
	PoolCache = append(PoolCache, block.KeyValue{
		Key:   data.Key,
		Value: data.Value,
	})
}

func pushBlock() {
	err := consensus.Push(PoolCache)
	if err != nil {
		log.Fatalln(err)
		return
	}
	PoolCache = make([]block.KeyValue, 0)
}
