package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/smallyunet/tmb/block"
)

func Server() {
	http.HandleFunc("/", root)
	http.HandleFunc("/post", post)
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HTTP server are running."))
}

func post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	var m map[string]string
	err = json.Unmarshal(body, &m)
	if err != nil {
		write(w, "Error json data format.")
		return
	}
	for k, v := range m {
		block.DataCache = append(block.DataCache, block.KeyValue{
			Key:   k,
			Value: v,
		})
	}
	block.DataMsg <- len(m)
	log.Println("DataCache size:", len(block.DataCache))
	write(w, "Service accepted the data.")
}

func write(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatalln(err)
	}
}
