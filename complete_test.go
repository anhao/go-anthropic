package anthropic

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestClient_CreateComplete(t *testing.T) {
	config := DefaultConfig("")
	client := NewClientWithConfig(config)

	complete, err := client.CreateComplete(context.Background(), CompleteRequest{
		Model:             "claude-2",
		Prompt:            "\n\nHuman: Hello\n\nAssistant:",
		MaxTokensToSample: 200,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(complete)
}

func TestClient_CreateCompleteStream(t *testing.T) {
	config := DefaultConfig("")
	client := NewClientWithConfig(config)

	stream, err := client.CreateCompleteStream(context.Background(), CompleteRequest{
		Model:             "claude-2",
		Prompt:            "\n\nHuman: Hello\n\nAssistant:",
		MaxTokensToSample: 200,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stream.Close()
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Println(response.Completion)
	}
}
