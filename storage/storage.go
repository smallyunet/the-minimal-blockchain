package storage

import (
	"encoding/json"
	"errors"
	"github.com/smallyunet/tmb/block"
	"github.com/smallyunet/tmb/config"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/smallyunet/tmb/util"
)

var path string

func Init() {
	path = config.DefaultDataPath
	value, ok := util.GetEnvVar(config.DataPathFlag)
	if ok {
		path = value
	}
	_, err := os.ReadDir(path)
	if err != nil {
		err = os.Mkdir(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func Set(height uint64, block *block.Block) error {
	if height == 0 {
		return errors.New("height must larger then zero")
	}
	ph := height - uint64(1)
	pb, err := Get(ph)
	if err != nil {
		return err
	}
	phv, err := util.GetHashCode(pb)
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

func Get(height uint64) (*block.Block, error) {
	p := getFilePath(height)
	b, err := os.ReadFile(p)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var blockData block.Block
	err = json.Unmarshal(b, &blockData)
	if err != nil {
		return nil, err
	}
	return &blockData, nil
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

func generateNextBlock(payload []block.KeyValue) (*block.Block, error) {
	ph, err := GetHeight()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	height := ph + 1
	pb, err := Get(ph)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	phv, err := util.GetHashCode(pb)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	block := &block.Block{}
	block.Prev = phv
	block.Height = height
	block.Payload = payload
	return block, nil
}

func Add(payload []block.KeyValue) (*block.Block, error) {
	nextBlock, err := generateNextBlock(payload)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	p := getFilePath(nextBlock.Height)
	b, err := json.Marshal(nextBlock)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = os.WriteFile(p, b, 0644)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return nextBlock, nil
}

func AddGenesisBlock() error {
	blockData := &block.Block{}
	blockData.Prev = ""
	blockData.Height = 0
	blockData.Payload = []block.KeyValue{}
	p := getFilePath(0)
	b, err := json.Marshal(blockData)
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
