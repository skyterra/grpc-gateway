package main

import (
	"net/http"
)

func RouterHandler(w http.ResponseWriter, r *http.Request) {
	GrpcRouter(w, r)
}

func main() {
	http.HandleFunc("/", RouterHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
