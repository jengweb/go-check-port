package main

import (
	"net/http"
)

func Response(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world JFrog."))
}

func main() {
	http.HandleFunc("/", Response)
	http.ListenAndServe(":8082", nil)
}
