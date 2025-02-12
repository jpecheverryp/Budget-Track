package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "HTTP Network Port")
    flag.Parse()

	mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./ui/static/"))

    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", getIndex)
    mux.HandleFunc("GET /transactions/view/{id}", getTransactionView)
    mux.HandleFunc("GET /transactions/create", getTransactionCreate)
    mux.HandleFunc("POST /transactions/create", postTransactionCreate)


	log.Printf("starting server on :%d", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), mux)
	log.Fatal(err)
}
