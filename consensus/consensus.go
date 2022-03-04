package consensus

import (
	"encoding/json"
	"github.com/smallyunet/tmb/client/tcp"
	"github.com/smallyunet/tmb/storage"
	"log"
)

func Push(payload string) error {
	// save in local first
	block, err := storage.Add(payload)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	// broadcast to other nodes
	b, err := json.Marshal(block)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	tcp.SendToAll(string(b))
	return nil
}

func HandleBlock(block *storage.Block) {
	// https://smallyu-net.translate.goog/2021/10/29/%E4%B8%80%E7%A7%8D%E5%9F%BA%E4%BA%8E%E2%80%9C%E8%87%AA%E6%88%91%E4%B8%AD%E5%BF%83%E4%B8%BB%E4%B9%89%E2%80%9D%E7%9A%84%E5%85%B1%E8%AF%86%E6%9C%BA%E5%88%B6/?_x_tr_sch=http&_x_tr_sl=auto&_x_tr_tl=en&_x_tr_hl=en-US&_x_tr_pto=nui
	// TODO handle this case
	log.Println("Receive: ", block)
}
