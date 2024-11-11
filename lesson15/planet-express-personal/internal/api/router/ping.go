package router

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponsePingJSON struct {
	Ping string `json:"ping"`
}

func GetPingHandler(w http.ResponseWriter, r *http.Request) {
	respJson := ResponsePingJSON{
		Ping: "pong",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(respJson)
	if err != nil {
		log.Fatal(err)
	}
}
