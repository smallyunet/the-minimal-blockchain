package hash

import (
	"strconv"

	"github.com/smallyunet/tmb/block"
)

func GetHashCode(block *block.Block) (string, error) {
	// TODO
	s := strconv.Itoa(int(block.Height))
	return s, nil
}
