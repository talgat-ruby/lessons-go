package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const PORT = 8080

type RequestJSON struct {
	Name string `json:"name"`
	Age  *int   `json:"age,omitempty"`
}

type ResponseJSON struct {
	Message string `json:"message"`
}

func serveHelloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("Hello World")))
	if err != nil {
		log.Fatal(err)
	}
}

func serveGoodBye(w http.ResponseWriter, r *http.Request) {
	reqJson := RequestJSON{}

	if r.Header.Get("Content-Type") == "application/json" {
		if err := json.NewDecoder(r.Body).Decode(&reqJson); err != nil {
			log.Fatal(err)
		}
	}

	mes := fmt.Sprintf("GoodBye %s! Your age is %d.", reqJson.Name, *reqJson.Age)

	if r.Header.Get("Accept") == "application/json" {
		respJson := ResponseJSON{
			Message: mes,
		}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(respJson)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	_, err := w.Write([]byte(mes))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx := context.Background()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", serveHelloWorld)
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
