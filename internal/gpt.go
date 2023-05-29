package internal

import (
	"context"
	"errors"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"io"
)

type GPT struct {
	client *openai.Client
	req    openai.ChatCompletionRequest
}

func (this *GPT) Init(role string) {
	this.client = openai.NewClient("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")
	this.req = openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 3000,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: role,
			},
		},
		Stream: false,
	}

}

func (this *GPT) Talk(msg string) {
	this.req.Messages = append(this.req.Messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})

	stream, err := this.client.CreateChatCompletionStream(context.Background(), this.req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			break
		}

		for _, choice := range response.Choices {
			print(choice.Delta.Content)
		}
	}

	println()
}
