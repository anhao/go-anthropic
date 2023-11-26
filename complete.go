package anthropic

import (
	"context"
	"errors"
	"net/http"
)

var (
	ErrCompleteStreamNotSupported        = errors.New("streaming is not supported with this method, please use CreateCompletionStream") //nolint:lll
	ErrCompletePromptNotEmpty            = errors.New("prompt is not empty")
	ErrCompleteMaxTokensToSmapleNotEmpty = errors.New("max_tokens_to_sample is not empty")
)

const (
	ClaudeV2Dot1        = "claude-2.1"
	ClaudeV2            = "claude-2"
	ClaudeV2Dot0        = "claude-2.0"
	ClaudeInstantV1     = "claude-instant-v1"
	ClaudeInstantV1Dot2 = "claude-instant-1.2"
)

type CompleteMetaData struct {
	UserId string `json:"user_id,omitempty"`
}
type CompleteRequest struct {
	Model             string           `json:"model"`
	Prompt            string           `json:"prompt"`
	MaxTokensToSample int              `json:"max_tokens_to_sample"`
	StopSequences     []string         `json:"stop_sequences,omitempty"`
	Temperature       float32          `json:"temperature,omitempty"`
	TopP              float32          `json:"top_p,omitempty"`
	TopK              int              `json:"top_k,omitempty"`
	Metadata          CompleteMetaData `json:"metadata,omitempty"`
	Stream            bool             `json:"stream,omitempty"`
}

type CompleteResponse struct {
	Completion string  `json:"completion"`
	StopReason *string `json:"stop_reason"`
	Model      string  `json:"model"`
	Stop       *string `json:"stop"`
	LogId      string  `json:"log_id"`
}

func (c *Client) CreateComplete(ctx context.Context, request CompleteRequest) (response CompleteResponse, err error) {
	if request.Stream {
		err = ErrCompleteStreamNotSupported
		return
	}
	if len(request.Prompt) == 0 {
		err = ErrCompletePromptNotEmpty
		return
	}
	if request.MaxTokensToSample == 0 {
		err = ErrCompleteMaxTokensToSmapleNotEmpty
		return
	}

	urlSuffix := "/complete"
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(urlSuffix, request.Model), withBody(request))
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
