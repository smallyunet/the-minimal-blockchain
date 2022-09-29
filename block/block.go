package block

import (
	"fmt"
	"strconv"
)

type Block struct {
	Prev    string     `json:"prev"`
	Height  uint64     `json:"height"`
	Payload []KeyValue `json:"payload"`
}

func (b Block) Serialize() string {
	h := strconv.Itoa(int(b.Height))
	return fmt.Sprintf("%s_%s_%s", b.Prev, h, b.Payload)
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
