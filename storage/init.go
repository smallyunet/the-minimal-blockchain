package storage

import (
	"os"

	"github.com/smallyunet/tmb/config"
	"github.com/smallyunet/tmb/util"
)

var path string

func init() {
	path = config.DefaultDataPath
	injectEnvVar()
	initDir()
}

func injectEnvVar() {
	value, ok := util.GetEnvVar("DataPath")
	if ok {
		path = value
	}
}

func initDir() {
	_, err := os.ReadDir(path)
	if err != nil {
		err = os.Mkdir(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}
