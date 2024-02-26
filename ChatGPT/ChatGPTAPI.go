package ChatGPT

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

const apiKey = "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"

func Callgpt(q string) string {
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: q,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "接口调用失败"
	}
	return resp.Choices[0].Message.Content
}
