package consensus

import (
	"encoding/json"
	"log"

	"github.com/smallyunet/tmb/block"
	"github.com/smallyunet/tmb/client/tcp"
	"github.com/smallyunet/tmb/storage"
)

func Push(payload []block.KeyValue) error {
	// save in local first
	blockData, err := storage.Add(payload)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	// broadcast to other nodes
	b, err := json.Marshal(blockData)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	tcp.SendToAll(string(b))
	return nil
}

func HandleBlock(block *block.Block) {
	// TODO handle this case
	log.Println("Receive: ", block)
}
