package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const PORT = 8080

func serveHelloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Fatal(err)
	}
}

func serveGoodBye(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("GoodBye"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", serveHelloWorld)
	mux.HandleFunc("/goodbye", serveGoodBye)

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", PORT),
		Handler:     mux,
		IdleTimeout: 2 * time.Minute,
	}

	log.Println("Listening on", PORT)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
