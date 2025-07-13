package service

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

// LLMClient 封装了 OpenAI 客户端
type LLMClient struct {
	client *openai.Client
}

// NewLLMClient 创建 LLMClient 实例，传入 API Key
func NewLLMClient(apiKey string) *LLMClient {
	return &LLMClient{
		client: openai.NewClient(apiKey),
	}
}

// GeneratePromptResponse 使用 GPT-3.5 Turbo 生成对话回复
func (c *LLMClient) GeneratePromptResponse(systemPrompt, userPrompt string) (string, error) {
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userPrompt,
				},
			},
		},
	)
	if err != nil {
		return "", fmt.Errorf("调用 OpenAI 失败: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("OpenAI 返回空响应")
	}

	return resp.Choices[0].Message.Content, nil
}
