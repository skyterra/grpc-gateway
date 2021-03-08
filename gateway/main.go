package main

import (
	"log"
	"net/http"
)

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	GrpcRouter(w, r)
}

func main() {
	log.Println("http server start")
	http.HandleFunc("/", RouterHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
