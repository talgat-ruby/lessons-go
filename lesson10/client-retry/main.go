package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	URL        = "https://jsonplaceholder.typicode.com/posts"
	MaxRetries = 3
)

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func fetchPost(ctx context.Context, id int) (*Post, error) {
	url := fmt.Sprintf("%s/%d", URL, id)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{Timeout: 5 * time.Second}

	var body []byte

	for range MaxRetries {
		resp, err := client.Do(req)
		if err != nil {
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			time.Sleep(time.Second)
			continue
		}

		body, err = func() ([]byte, error) {
			defer resp.Body.Close()
			return io.ReadAll(resp.Body)
		}()
		if err == nil {
			break
		}
	}

	if body == nil {
		return nil, fmt.Errorf("failed to fetch post after retries")
	}

	var post Post
	if err := json.Unmarshal(body, &post); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return &post, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	postIDs := []int{1, 2, 3, 4, 5}
	posts := make([]*Post, 0, len(postIDs))
	var wg sync.WaitGroup

	for _, id := range postIDs {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			post, err := fetchPost(ctx, id)
			if err != nil {
				fmt.Printf("Error fetching post %d: %v\n", id, err)
				return
			}
			posts = append(posts, post)
		}(id)
	}

	wg.Wait()

	for _, post := range posts {
		if post != nil {
			fmt.Printf("Post ID: %d, Title: %s\n", post.ID, post.Title)
		}
	}
}
