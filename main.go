package main

import (
	"log"
	"net/http"

	"web/server"
)

func main() {
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.HandleFunc("/", server.PageRender)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
