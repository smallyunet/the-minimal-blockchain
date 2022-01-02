package block

import (
	"encoding/json"
	"log"
	"time"

	"github.com/smallyunet/tmb/consensus"
	"github.com/smallyunet/tmb/storage"
)

var DataCache = make([]storage.KeyValue, 0)

var DataMsg = make(chan int, 1)

func init() {
	go Accept()
}

func Accept() {
	ticker := time.NewTicker(time.Duration(blockTime) * time.Millisecond)

	for {
		select {
		case <-DataMsg:
			if uint64(len(DataCache)) >= blockSize {
				pushData()
			}
		case <-ticker.C:
			pushData()
		default:
			// do nothing
		}
	}
}

func pushData() {
	d, err := json.Marshal(DataCache)
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = consensus.Push(string(d))
	if err != nil {
		log.Fatalln(err)
		return
	}
	DataCache = make([]storage.KeyValue, 0)
}
