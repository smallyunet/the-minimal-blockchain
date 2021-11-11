package http

import (
	"log"
	"net/http"
)

func Server() {
	http.HandleFunc("/", root)
	err := http.ListenAndServe(":25001", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service are running."))
}
