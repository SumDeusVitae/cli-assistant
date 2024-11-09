package gpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GptReply struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
			Refusal any    `json:"refusal"`
		} `json:"message"`
		Logprobs     any    `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens        int `json:"prompt_tokens"`
		CompletionTokens    int `json:"completion_tokens"`
		TotalTokens         int `json:"total_tokens"`
		PromptTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
			AudioTokens  int `json:"audio_tokens"`
		} `json:"prompt_tokens_details"`
		CompletionTokensDetails struct {
			ReasoningTokens          int `json:"reasoning_tokens"`
			AudioTokens              int `json:"audio_tokens"`
			AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
			RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
		} `json:"completion_tokens_details"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

func RequestGPT(reqMessage, apiKey string) (string, error) {
	apiUrl := "https://api.openai.com/v1/chat/completions"
	if apiKey == "" {
		return "", errors.New("OpenAI API key not provided 'OPENAI_API_KEY'")
	}
	type reqStruct struct {
		Model    string `json:"model"`
		Messages []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"messages"`
		Temperature float64 `json:"temperature"`
	}
	request := reqStruct{
		Model:       "gpt-4o-mini",
		Temperature: 0.7,
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "user", Content: reqMessage},
		},
	}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("error creating JSON payload: %v", err)
	}

	// Create HTTP REQUEST
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response
	dat, err := io.ReadAll(resp.Body)
	if err != nil {

		return "", fmt.Errorf("error reading response:%v", err)
	}
	response := GptReply{}
	err = json.Unmarshal(dat, &response)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response:%v", err)
	}
	choices := response.Choices
	// log.Printf("choices: %v", choices)
	// log.Printf("choices[0].Message.Content?: %v", choices[0].Message.Content)

	if len(choices) > 0 && choices[0].Message.Content != "" {
		return choices[0].Message.Content, nil
	}
	return "No reply received from GPT", nil

}
