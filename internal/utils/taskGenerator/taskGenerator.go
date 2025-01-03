package taskGenerator

import (
	"fmt"
	"gera-ai/internal/utils/openai"
	"strings"
)

const (
	maxOpenAITokens = 100
	openAIModel     = "gpt-3.5-turbo"
)

// TaskGenerator предоставляет функции для генерации и анализа задач
type TaskGenerator struct {
	client *openai.Client
}

// NewTaskGenerator создает новый TaskGenerator
func NewTaskGenerator(apiKey, proxyURL string) (*TaskGenerator, error) {
	client, err := openai.NewClient(&openai.Config{
		APIKey:   apiKey,
		ProxyURL: proxyURL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create OpenAI client: %w", err)
	}

	return &TaskGenerator{client: client}, nil
}

// GenerateTaskWithInterests генерирует задачу с учетом интересов
func (tg *TaskGenerator) GenerateTaskWithInterests(condition string, interests []string) (string, error) {
	// Формируем запрос с учетом интересов
	prompt := fmt.Sprintf(
		`условие: %s
добавь в это условие сюжет по следующим интересам: %s.
Сделай сюжет максимально связным с интересами человека, добавь туда юмор, можешь черный, чтобы ему было интересно решать задачу.
Не усложняй условие. Не меняй значения в условии.
Не пиши ответ и не обьясняй задачу. Скинь новую задачу с добавленным в нее сюжетом. Увеличивай ее ДО 1.5x, где x - исходный размер задачи.
Размер задачи должен быть не больше 1000 символов.
Выдай только текст нового условия.`,
		condition, strings.Join(interests, "; "),
	)

	// Вызываем OpenAI API с подготовленным запросом
	response, err := tg.client.CallOpenAI(prompt, openAIModel, maxOpenAITokens)
	if err != nil {
		return "", fmt.Errorf("failed to generate task with interests: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	// Возвращаем сгенерированный текст
	return response.Choices[0].Message.Content, nil
}

// GenerateTaskWithNoInterests генерирует задачу с "реалистичным" сюжетом
func (tg *TaskGenerator) GenerateTaskWithNoInterests(condition string) (string, error) {
	// Формируем запрос с "реалистичным" сюжетом
	prompt := fmt.Sprintf(
		`условие: %s
Добавь сюжет в задачу, сделай ее максимально приближенной к реальной, суровой жизни Русских. Не меняй ответ на задачу.
Не пиши ответ и не обьясняй задачу.
Скинь новую задачу с добавленным в нее сюжетом. Увеличивай ее ДО 1.5x, где x - исходный размер задачи.
Размер задачи должен быть не больше 1000 символов.
Выдай только текст нового условия.`,
		condition,
	)

	// Вызываем OpenAI API с подготовленным запросом
	response, err := tg.client.CallOpenAI(prompt, openAIModel, maxOpenAITokens)
	if err != nil {
		return "", fmt.Errorf("failed to generate task with life plot: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	// Возвращаем сгенерированный текст
	return response.Choices[0].Message.Content, nil
}

// GenerateAnswer делает разбор задачи
func (tg *TaskGenerator) GenerateAnswer(condition string) (string, error) {
	// Формируем запрос для анализа задачи
	prompt := fmt.Sprintf(
		`условие: %s
Сделай разбор задачи. Раскрой весь сюжет. Покажи формулы в этой задаче и темы, на которые нацелена эта задача.
Добавь туда юмор.
Размер разбора должен быть не больше 100 символов.
Выдай только текст разбора задачи. Формулы пиши обычным текстом.`,
		condition,
	)

	// Вызываем OpenAI API с подготовленным запросом
	response, err := tg.client.CallOpenAI(prompt, openAIModel, maxOpenAITokens)
	if err != nil {
		return "", fmt.Errorf("failed to analyze task: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	// Возвращаем сгенерированный текст
	return response.Choices[0].Message.Content, nil
}
