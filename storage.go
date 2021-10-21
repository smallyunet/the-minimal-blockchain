package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var path string

func init() {
	path = "./data/"
}

func Set(height uint64) error {
	// TODO end with seperate
	path += strconv.Itoa(int(height)) + ".json"
	os.WriteFile(path, []byte("abc"), 0644)
	return nil
}

func Get(height uint64) (*Block, error) {
	return nil, nil
}

func GetHeight() (uint64, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return uint64(len(files)), nil
}
