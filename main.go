package main

import (
	"net/http"
)

func helloWord(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", helloWord)
	http.ListenAndServe(":8080", nil)
}
