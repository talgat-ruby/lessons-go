package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const URL = "https://jsonplaceholder.typicode.com/posts/1"

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	response, err := http.Get(URL)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var post Post

	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	fmt.Printf("Post ID: %d\n", post.ID)
	fmt.Printf("User ID: %d\n", post.UserID)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Body: %s\n", post.Body)
}
