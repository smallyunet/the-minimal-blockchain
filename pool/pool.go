package pool

import (
	"encoding/json"
	"log"
	"time"

	"github.com/smallyunet/tmb/consensus"
	"github.com/smallyunet/tmb/storage"
)

var PoolCache = make([]storage.KeyValue, 0)

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

func PushToPool(data storage.KeyValue) {
	PoolCache = append(PoolCache, storage.KeyValue{
		Key:   data.Key,
		Value: data.Value,
	})
}

func pushBlock() {
	d, err := json.Marshal(PoolCache)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = consensus.Push(string(d))
	if err != nil {
		log.Fatalln(err)
		return
	}
	PoolCache = make([]storage.KeyValue, 0)
}
