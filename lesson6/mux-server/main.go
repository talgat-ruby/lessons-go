package main

import (
	"fmt"
	"log"
	"net/http"
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
	http.HandleFunc("/hello", serveHelloWorld)
	http.HandleFunc("/goodbye", serveGoodBye)
	log.Println("Listening on", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatal(err)
	}
}
