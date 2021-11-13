package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", root)
	http.HandleFunc("/post", post)
	err := http.ListenAndServe(":25001", nil)
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

	w.Write([]byte("Service accepted the data."))
}
