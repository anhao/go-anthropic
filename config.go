package anthropic

import "net/http"

const (
	anthropicAPIURLv1              = "https://api.anthropic.com/v1"
	defaultEmptyMessagesLimit uint = 300

	anthropicVersion = "2023-06-01"
)

type ClientConfig struct {
	ApiKey             string
	Version            string
	HTTPClient         *http.Client
	EmptyMessagesLimit uint
	BaseURL            string
}

func DefaultConfig(apikey string) ClientConfig {
	return ClientConfig{
		ApiKey:             apikey,
		Version:            anthropicVersion,
		HTTPClient:         &http.Client{},
		EmptyMessagesLimit: defaultEmptyMessagesLimit,
		BaseURL:            anthropicAPIURLv1,
	}
}
