package main

import (
	"log"
	http "net/http"
)

const (
	Server = "localhost"
	Port   = "8080"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World</h1>"))
}

func StartServer() {
	srv := &http.Server{
		Addr: Server + ":" + Port,
	}
	http.HandleFunc("/", HelloWorld)
	log.Fatal(srv.ListenAndServe())
}

func main() {
	StartServer()
}
