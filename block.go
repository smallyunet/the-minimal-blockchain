package main

type Block struct {
	Prev    string `json:"prev"`
	Height  uint64 `json:"Height"`
	Payload string `json:"Payload"`
}

func init() {

}
