package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/smallyunet/tmb/block"
)

func GetHashCode(b *block.Block) (string, error) {
	s := b.Serialize()
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:]), nil
}
