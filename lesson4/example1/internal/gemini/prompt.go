package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const mainURL string = "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent"

func (ai *Gemini) SendPrompt(prompt string) (string, error) {
	req, err := createRequest(ai.apiKey, prompt)
	if err != nil {
		slog.Error("could not create Request object")
		return "", err
	}

	resp, err := makeRequest(req)
	if err != nil {
		slog.Error("could not make Request")
		return "", err
	}

	text, err := getTextFromResponse(resp)
	if err != nil {
		slog.Error("could not get text from response")
		return "", err
	}

	return text, nil
}

type Part struct {
	Text string `json:"text"`
}

type Content struct {
	Parts []*Part `json:"parts"`
}

type GeminiRequestBody struct {
	Contents []*Content `json:"contents"`
}

func createRequest(apiKey string, prompt string) (*http.Request, error) {
	requestURL := fmt.Sprintf("%s?key=%s", mainURL, apiKey)

	body := GeminiRequestBody{
		Contents: []*Content{
			{
				Parts: []*Part{
					{
						Text: fmt.Sprintf("%s. Please send it in LaTEX", promqpt),
					},
				},
			},
		},
	}

	j, err := json.Marshal(body)
	if err != nil {
		slog.Error("client: could not marshal json: %s\n", err)
		return nil, err
	}

	reqBody := bytes.NewBuffer(j)

	req, err := http.NewRequest(http.MethodPost, requestURL, reqBody)
	if err != nil {
		slog.Error("client: could not create request: %s\n", err)
		return nil, err
	}

	return req, nil
}

func makeRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("client: error making http request: %s\n", err)
		return nil, err
	}

	return resp, nil
}

type Candidates struct {
	Content *Content `json:"content"`
}

type GeminiResponseBody struct {
	Candidates []*Candidates `json:"candidates"`
}

func getTextFromResponse(resp *http.Response) (string, error) {
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("read body error")
		return "", err
	}
	defer resp.Body.Close()

	var responseObject GeminiResponseBody
	if err := json.Unmarshal(responseData, &responseObject); err != nil {
		slog.Error("unmarshal error")
		return "", err
	}

	if len(responseObject.Candidates) <= 0 {
		return "", err
	}

	if responseObject.Candidates[0].Content == nil {
		return "", err
	}

	if len(responseObject.Candidates[0].Content.Parts) <= 0 {
		return "", err
	}

	return responseObject.Candidates[0].Content.Parts[0].Text, err
}
