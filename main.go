package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/put", HandlePut)
	http.HandleFunc("/get", HandleGet)
	http.HandleFunc("/delete", HandleDelete)
	http.HandleFunc("/list", HandleList)
	http.ListenAndServe(":8080", nil)
}