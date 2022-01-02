package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/smallyunet/tmb/block"
	"github.com/smallyunet/tmb/storage"
)

func Server() {
	http.HandleFunc("/", root)
	err := http.ListenAndServe(":"+httpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte("HTTP server are running."))
	case "/post":
		post(w, r)
	default:
	}

	if strings.HasPrefix(r.URL.Path, "/get") {
		get(w, r)
	}
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

func get(w http.ResponseWriter, r *http.Request) {
	p := strings.Split(r.URL.Path, "/")
	if len(p) < 3 {
		write(w, "Error request path.")
		return
	}
	h, err := strconv.Atoi(p[2])
	if err != nil {
		write(w, "Error request path.")
		return
	}
	height, err := storage.GetHeight()
	if err != nil {
		write(w, "Error request path.")
		return
	}
	if h > int(height) {
		write(w, "No data.")
		return
	}
	block, err := storage.Get(uint64(h))
	if err != nil {
		write(w, "Error request path.")
		return
	}
	s, err := json.Marshal(block)
	if err != nil {
		write(w, "Error json data format.")
		return
	}
	write(w, string(s))
}

func write(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatalln(err)
	}
}
