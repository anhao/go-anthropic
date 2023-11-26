# Go Anthropic
[![Go Reference](https://pkg.go.dev/badge/github.com/anhao/go-anthropic.svg)](https://pkg.go.dev/github.com/anhao/go-anthropic)

This library provides unofficial Go clients for [Anthropic API](https://console.anthropic.com/). We support: 

* Complete


## Installation

```
go get github.com/anhao/go-anthropic
```
Currently, go-anthropic requires Go version 1.18 or greater.


## Usage

### Anthropic example usage:

```go
package main

import (
	"context"
	"fmt"
	anthropic "github.com/anhao/go-anthropic"
)

func main() {
	client := anthropic.NewClient("your token")
	resp, err := client.CreateComplete(
		context.Background(),
		anthropic.CompleteRequest{
			Model: anthropic.ClaudeV2Dot1,
			Prompt: anthropic.GetPromptFromString("Hello"),
			MaxTokensToSample: 400,
		},
	)

	if err != nil {
		fmt.Printf("Complete error: %v\n", err)
		return
	}

	fmt.Println(resp.Completion)
}

```

### Getting an Anthropic API Key:

1. Visit the OpenAI website at [https://console.anthropic.com/account/keys](https://console.anthropic.com/account/keys).
2. If you don't have an account, click on "Sign Up" to create one. If you do, click "Log In".
3. Once logged in, navigate to your API key management page.
4. Click on "Create Key".
5. Enter a name for your new key, then click "Create secret key".
6. Your new API key will be displayed. Use this key to interact with the Anthropic API.

**Note:** Your API key is sensitive information. Do not share it with anyone.

### Other examples:

<details>
<summary>Completion streaming </summary>

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	anthropic "github.com/anhao/go-anthropic"
)

func main() {
	c := anthropic.NewClient("your token")
	ctx := context.Background()

	req := anthropic.CompleteRequest{
		Model:     anthropic.ClaudeV2,
		Stream: true,
		MaxTokensToSample: 400,
		Prompt: anthropic.GetPromptFromString("Hello"),
	}
	stream, err := c.CreateCompleteStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
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

		fmt.Printf(response.Completion)
	}
}
```
</details>






## Thank you
- [go-openai](https://github.com/sashabaranov/go-openai)