package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const apiURL = "https://api.openai.com/v1/chat/completions"

// Config содержит настройки для общения с OpenAI API
type Config struct {
	APIKey   string
	ProxyURL string // Адрес прокси (если нужен)
}

// RequestBody структура для отправки данных в OpenAI API
type RequestBody struct {
	Model     string              `json:"model"`
	Messages  []map[string]string `json:"messages"`
	MaxTokens int                 `json:"max_tokens"`
}

// ResponseBody структура для получения ответа от OpenAI API
type ResponseBody struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// Client реализует клиента для общения с OpenAI API
type Client struct {
	config *Config
	client *http.Client
}

// NewClient создает нового клиента для OpenAI API
func NewClient(config *Config) (*Client, error) {
	transport := &http.Transport{}

	if config.ProxyURL != "" {
		proxy, err := url.Parse(config.ProxyURL)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy URL: %w", err)
		}
		transport.Proxy = http.ProxyURL(proxy)
	}

	httpClient := &http.Client{Transport: transport}

	return &Client{
		config: config,
		client: httpClient,
	}, nil
}

// CallOpenAI отправляет запрос к OpenAI API
func (c *Client) CallOpenAI(prompt string, model string, maxTokens int) (ResponseBody, error) {
	// Формируем запрос с правильной структурой
	requestBody := RequestBody{
		Model:     model,
		Messages:  []map[string]string{{"role": "user", "content": prompt}},
		MaxTokens: maxTokens,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return ResponseBody{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return ResponseBody{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	req.Header.Set("Content-Type", "application/json")

	// Отправка запроса
	resp, err := c.client.Do(req)
	if err != nil {
		return ResponseBody{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Проверка на успешный статус
	if resp.StatusCode != http.StatusOK {
		return ResponseBody{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Чтение и разбор ответа
	var response ResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return ResponseBody{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}
