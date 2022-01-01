package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/smallyunet/tmb/block"
	"github.com/smallyunet/tmb/hash"
)

func Set(height uint64, block *block.Block) error {
	if height == 0 {
		return errors.New("height must larger then zero")
	}
	ph := height - uint64(1)
	pb, err := Get(ph)
	if err != nil {
		return err
	}
	phv, err := hash.GetHashCode(pb)
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
	err = os.WriteFile(p, []byte(b), 0644)
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
	var block block.Block
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
	// TODO filter end without .json
	return uint64(len(files)), nil
}

func Add(block *block.Block) error {
	ph, err := GetHeight()
	if err != nil {
		return err
	}
	height := ph + 1
	pb, err := Get(ph)
	if err != nil {
		return err
	}
	phv, err := hash.GetHashCode(pb)
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
	err = os.WriteFile(p, []byte(b), 0644)
	if err != nil {
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
