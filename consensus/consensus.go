package consensus

import (
	"github.com/smallyunet/tmb/client/tcp"
	"log"

	"github.com/smallyunet/tmb/storage"
)

func Push(payload string) error {
	err := storage.Add(payload)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	// broadcast to other nodes
	// TODO all block content ranther than payload
	tcp.SendToAll(payload)
	return nil
}
