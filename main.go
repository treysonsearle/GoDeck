package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/deal", dealCardHandler)
	http.HandleFunc("/shuffle", shuffleHandler)
	http.HandleFunc("/discard", discardHandler)
	http.HandleFunc("/cut", cutHandler)
	http.HandleFunc("/order", orderHandler)
	http.HandleFunc("/rebuild", rebuildHandler)
	http.HandleFunc("/cheat", cheatHandler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
