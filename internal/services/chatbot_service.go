package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"hris_backend/internal/request"
	"hris_backend/internal/response"
	"hris_backend/pkg"
)

type ChatbotServices interface {
	Ask(req request.ChatbotRequest) (response.ChatbotResponse, error)
}

type chatbotServices struct {
	apiKey    string
	client    *http.Client
	knowledge string
}

func NewChatbotService(apiKey string, companyFilePath string) ChatbotServices {
	knowledge := ""
	if companyFilePath != "" {
		b, err := os.ReadFile(companyFilePath)
		if err == nil {
			knowledge = string(b)
		} else {
			fmt.Printf("⚠️ gagal baca knowledge file: %v\n", err)
		}
	}

	return &chatbotServices{
		apiKey:    apiKey,
		client:    &http.Client{},
		knowledge: knowledge,
	}
}

func (s *chatbotServices) Ask(req request.ChatbotRequest) (response.ChatbotResponse, error) {

	if err := pkg.Validate.Struct(req); err != nil {
		return response.ChatbotResponse{}, err
	}

	messages := []map[string]string{}

	if s.knowledge != "" {

		messages = append(messages, map[string]string{
			"role": "system",
			"content": "Kamu adalah asisten perusahaan. " +
				"Gunakan informasi berikut tentang perusahaan untuk menjawab pertanyaan: \n\n" +
				s.knowledge +
				"\n\nJika pertanyaan tidak relevan dengan informasi ini, jawab dengan jawaban umum.",
		})
	}

	messages = append(messages, map[string]string{
		"role":    "user",
		"content": req.Message,
	})

	payload := map[string]interface{}{
		"model":    req.Model,
		"messages": messages,
	}

	body, _ := json.Marshal(payload)

	httpReq, err := http.NewRequest(
		"POST",
		"https://openrouter.ai/api/v1/chat/completions",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return response.ChatbotResponse{}, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	resp, err := s.client.Do(httpReq)
	if err != nil {
		return response.ChatbotResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(resp.Body)
		return response.ChatbotResponse{}, errors.New(string(b))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return response.ChatbotResponse{}, err
	}

	if len(result.Choices) == 0 {
		return response.ChatbotResponse{}, errors.New("no reply from chatbot")
	}

	return response.ChatbotResponse{
		Reply: result.Choices[0].Message.Content,
	}, nil
}
