package http

import (
	"encoding/json"
	"github.com/smallyunet/tmb/block"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/smallyunet/tmb/pool"
	"github.com/smallyunet/tmb/storage"
)

func Server(wg *sync.WaitGroup) {
	http.HandleFunc("/", root)
	failed := make(chan bool, 1)
	go func() {
		err := http.ListenAndServe(":"+httpPort, nil)
		if err != nil {
			failed <- true
			log.Fatal(err)
		}
	}()
	log.Println("HTTP server running on port", httpPort)
	wg.Done()
	if <-failed {
		log.Fatal("HTTP server stopped.")
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/post") {
		post(w, r)
	}
	if strings.HasPrefix(r.URL.Path, "/get") {
		get(w, r)
	}
	if strings.HasPrefix(r.URL.Path, "/info") {
		info(w, r)
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var m map[string]string
	err = json.Unmarshal(body, &m)
	if err != nil {
		write(w, "Error json data format.")
		return
	}
	for k, v := range m {
		pool.PushTxToPool(block.KeyValue{
			Key:   k,
			Value: v,
		})
	}
	pool.DataMsg <- len(m)
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
	b, err := storage.Get(uint64(h))
	if err != nil {
		write(w, "Error request path.")
		return
	}
	s, err := json.Marshal(b)
	if err != nil {
		write(w, "Error json data format.")
		return
	}
	write(w, string(s))
}

func info(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]interface{})
	height, err := storage.GetHeight()
	if err != nil {
		write(w, "Error request path.")
		return
	}
	res["height"] = height
	resBytes, err := json.Marshal(res)
	if err != nil {
		write(w, "Error json data format.")
		return
	}
	write(w, string(resBytes))
}

func write(w http.ResponseWriter, msg string) {
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatalln(err)
	}
}
