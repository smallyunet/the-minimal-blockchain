package hash

import (
	"strconv"
	"the-minimal-blockchain/block"
)

func GetHashCode(block *block.Block) (string, error) {
	// TODO
	s := strconv.Itoa(int(block.Height))
	return s, nil
}
