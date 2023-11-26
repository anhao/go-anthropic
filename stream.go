package anthropic

import (
	"errors"
)

var (
	ErrTooManyEmptyStreamMessages = errors.New("stream has sent too many empty messages")
)

type CompletionStream struct {
	*streamReader[CompleteResponse]
}
