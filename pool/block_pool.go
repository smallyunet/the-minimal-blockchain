package pool

import (
	"github.com/smallyunet/tmb/block"
	"github.com/smallyunet/tmb/consensus"
	"github.com/smallyunet/tmb/storage"
	"log"
	"time"
)

var BlockPool = make([]*block.Block, 0)

var BlockMsg = make(chan int, 1)

func AcceptBlock() {
	ticker := time.NewTicker(time.Duration(blockTime) * time.Millisecond)

	for {
		select {
		case <-BlockMsg:
			// TODO set this condition
			//if len(BlockPool) >= len(tcp.RouteTable) {
			//	pushBlock()
			//}
		case <-ticker.C:
			if len(BlockPool) > 0 {
				pushBlock()
			}
		default:
			// do nothing
		}
	}
}

func PushBlockToPool(data *block.Block) {
	height, err := storage.GetHeight()
	if err != nil {
		log.Println(err)
		return
	}
	if data.Height == height+1 {
		BlockPool = append(BlockPool, data)
	}
}

func pushBlock() {
	err := consensus.HandleBlockFromPool(BlockPool)
	if err != nil {
		log.Fatalln(err)
		return
	}
	BlockPool = make([]*block.Block, 0)
}
