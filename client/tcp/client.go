package tcp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/smallyunet/tmb/block"
	"github.com/smallyunet/tmb/pool"
	"github.com/smallyunet/tmb/storage"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func Send(protocol, address, msg string) string {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Fprint(conn, msg)
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	return buf.String()
}

func SendToAll(msg string) {
	for k, _ := range RouteTable {
		// TODO change the default protocol
		conn, err := net.Dial("tcp", k)
		if err != nil {
			log.Println("Broadcast block data:", err)
			return
		}
		defer conn.Close()
		fmt.Fprint(conn, msg)
		var buf bytes.Buffer
		io.Copy(&buf, conn)
	}
}

func GetLatestBlock() {
	ticker := time.NewTicker(time.Duration(blockTime) * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			height, err := storage.GetHeight()
			if err != nil {
				log.Println(err)
			}
			for k, _ := range RouteTable {
				requestURL := fmt.Sprintf("http://%s/get/%d", k, height+1)
				get, err := http.Get(requestURL)
				if err != nil {
					continue
				}
				var b *block.Block
				all, err := io.ReadAll(get.Body)
				if err != nil {
					continue
				}
				err = json.Unmarshal(all, &b)
				if err != nil {
					continue
				}
				pool.PushBlockToPool(b)
			}
		}
	}
}
