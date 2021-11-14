package block

import "log"

type Block struct {
	Prev    string `json:"prev"`
	Height  uint64 `json:"height"`
	Payload string `json:"payload"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

var DataCache = make([]KeyValue, 0)

var DataMsg = make(chan int, 1)

func init() {
	go Accept()
}

func Accept() {
	for {
		select {
		case <-DataMsg:
			log.Println("DataMsg", len(DataCache))
		default:
			
		}
	}
}
