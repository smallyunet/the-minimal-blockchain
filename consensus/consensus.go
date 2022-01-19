package consensus

import (
	"github.com/smallyunet/tmb/client"
	"log"

	"github.com/smallyunet/tmb/storage"
)

func Push(payload string) error {
	err := storage.Add(payload)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	// bordcast to other nodes
	// TODO all block content ranther than payload
	client.SendToAll(payload)
	return nil
}
