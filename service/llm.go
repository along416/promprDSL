package service

import (
	"context"
	"errors"
	"os"

	openai "github.com/sashabaranov/go-openai" // 这是一个Go OpenAI客户端库，你也可以换别的
)

// Client 是对外调用的大模型客户端封装
type Client struct {
	cli *openai.Client
}

// NewClient 初始化客户端
func NewClient() *Client {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY 环境变量未设置")
	}
	return &Client{
		cli: openai.NewClient(apiKey),
	}
}

// GenerateText 传入prompt，调用大模型生成文本
func (c *Client) GenerateText(ctx context.Context, prompt string) (string, error) {
	if prompt == "" {
		return "", errors.New("prompt 不能为空")
	}

	resp, err := c.cli.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: "gpt-4o-mini", // 换成你要用的模型
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	})
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", errors.New("没有生成内容")
	}

	return resp.Choices[0].Message.Content, nil
}
// // 提取 JSON 数组
// func extractJSONArray(text string) string {
// 	re := regexp.MustCompile(`(?s)(\[.*?\])`)
// 	matches := re.FindStringSubmatch(text)
// 	if len(matches) > 1 {
// 		return matches[1] // 第一个子匹配是数组
// 	}
// 	return ""
// }
