package consensus

import (
	"encoding/json"
	"errors"
	"github.com/smallyunet/tmb/util"
	"log"
	"math/rand"

	"github.com/smallyunet/tmb/block"
	//"github.com/smallyunet/tmb/client/tcp"
	"github.com/smallyunet/tmb/storage"
)

func SaveBlockByTx(payload []block.KeyValue) error {
	// save in local first
	blockData, err := storage.Add(payload)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	// broadcast to other nodes
	b, err := json.Marshal(blockData)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	//tcp.SendToAll(string(b))
	log.Println("Send to all nodes: " + string(b))
	return nil
}

func SaveBlockByBlock(block *block.Block) error {
	_, err := storage.Add(block.Payload)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func HandleBlockFromPool(pool []*block.Block) error {
	b := selectBlock(pool)
	r := verifyBlock(b)
	if r {
		_, err := storage.Add(b.Payload)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		return errors.New("verify block failed")
	}
	return nil
}

func selectBlock(pool []*block.Block) *block.Block {
	i := rand.Intn(len(pool))
	b := pool[i]
	return b
}

func verifyBlock(b *block.Block) bool {
	height, err := storage.GetHeight()
	if err != nil {
		return false
	}
	get, err := storage.Get(height)
	if err != nil {
		return false
	}
	code, err := util.GetHashCode(b)
	if err != nil {
		return false
	}
	if code != get.Prev {
		return false
	}
	return false
}
