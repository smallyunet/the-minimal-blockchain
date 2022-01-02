package consensus

import (
	"log"

	"github.com/smallyunet/tmb/storage"
)

func Push(payload string) error {
	// TODO consensus logic
	err := storage.Add(payload)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return nil
}
