package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/smallyunet/tmb/util"
)

type Block struct {
	Prev    string `json:"prev"`
	Height  uint64 `json:"height"`
	Payload string `json:"payload"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Set(height uint64, block *Block) error {
	if height == 0 {
		return errors.New("height must larger then zero")
	}
	ph := height - uint64(1)
	pb, err := Get(ph)
	if err != nil {
		return err
	}
	phv, err := util.GetHashCode(pb.Payload)
	if err != nil {
		return err
	}
	block.Prev = phv
	block.Height = height
	p := getFilePath(height)
	b, err := json.Marshal(block)
	if err != nil {
		return err
	}
	err = os.WriteFile(p, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Get(height uint64) (*Block, error) {
	p := getFilePath(height)
	b, err := os.ReadFile(p)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var block Block
	err = json.Unmarshal(b, &block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}

func GetHeight() (uint64, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	if len(files) == 0 {
		err := AddGenesisBlock()
		if err != nil {
			log.Fatalln(err)
		}
		return 0, nil
	}
	// TODO filter end without .json
	return uint64(len(files)) - 1, nil
}

func Add(payload string) error {
	ph, err := GetHeight()
	if err != nil {
		log.Fatalln(err)
		return err
	}
	height := ph + 1
	pb, err := Get(ph)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	phv, err := util.GetHashCode(pb.Payload)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	block := &Block{}
	block.Prev = phv
	block.Height = height
	block.Payload = payload
	p := getFilePath(height)
	b, err := json.Marshal(block)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	err = os.WriteFile(p, b, 0644)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func AddGenesisBlock() error {
	block := &Block{}
	block.Prev = ""
	block.Height = 0
	block.Payload = ""
	p := getFilePath(0)
	b, err := json.Marshal(block)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	err = os.WriteFile(p, b, 0644)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func getFilePath(height uint64) string {
	var p string
	if strings.HasSuffix(path, string(os.PathSeparator)) {
		p = path + strconv.Itoa(int(height)) + ".json"
	} else {
		p = path + string(os.PathSeparator) + strconv.Itoa(int(height)) + ".json"
	}
	return p
}
