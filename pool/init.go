package pool

import "github.com/smallyunet/tmb/config"

var blockSize uint64
var blockTime uint64

func init() {
	blockSize = config.DefaultBlockSize
	blockTime = config.DefaultBlockTime

	go AcceptTx()
	go AcceptBlock()
}
