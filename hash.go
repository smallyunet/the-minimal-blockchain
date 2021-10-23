package main

import "strconv"

func GetHashCode(block *Block) (string, error) {
	// TODO
	s := strconv.Itoa(int(block.Height))
	return s, nil
}
