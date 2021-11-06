package block

type Block struct {
	Prev    string `json:"prev"`
	Height  uint64 `json:"height"`
	Payload string `json:"payload"`
}

func init() {

}
