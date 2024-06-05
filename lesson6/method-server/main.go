package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const PORT = 8080

func serveHelloWorld(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	_, err := w.Write([]byte(fmt.Sprintf("Hello %s", name)))
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
	ctx := context.Background()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", serveHelloWorld)
	mux.HandleFunc("POST /goodbye", serveGoodBye)

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", PORT),
		Handler:     mux,
		IdleTimeout: 2 * time.Minute,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	log.Println("Listening on", PORT)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
