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
const TEXT_KEY = "text"

type handler struct {
	text string
}

func (h *handler) serveHelloWorld(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	text, ok := ctx.Value(TEXT_KEY).(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Printf("Our text from context::%s", text)

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
	ctx, cancel := context.WithCancel(context.Background())

	text := "From context"
	newCtx := context.WithValue(ctx, TEXT_KEY, text)

	h := &handler{text: text}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", h.serveHelloWorld)
	mux.HandleFunc("/goodbye", serveGoodBye)

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", PORT),
		Handler:     mux,
		IdleTimeout: 2 * time.Minute,
		BaseContext: func(_ net.Listener) context.Context {
			return newCtx
		},
	}

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()

	go func() {
		<-newCtx.Done()
		_ = srv.Shutdown(context.Background())
	}()

	log.Println("Listening on", PORT)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
