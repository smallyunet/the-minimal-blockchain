package block

type Block struct {
	Prev    string `json:"prev"`
	Height  uint64 `json:"height"`
	Payload string `json:"payload"`
}

var DataCache = make(map[string]string)

func init() {

}
