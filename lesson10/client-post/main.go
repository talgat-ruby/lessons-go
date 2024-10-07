package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const URL = "https://jsonplaceholder.typicode.com/posts"

type Post struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	newPost := Post{
		UserID: 1,
		Title:  "foo",
		Body:   "bar",
	}

	jsonData, err := json.Marshal(newPost)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))

	var responsePost Post
	err = json.Unmarshal(body, &responsePost)
	if err != nil {
		log.Fatalf("Error unmarshaling response: %v", err)
	}

	fmt.Printf("Created post with ID: %d\n", responsePost.UserID)
}
