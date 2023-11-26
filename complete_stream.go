package anthropic

import "context"

func (c *Client) CreateCompleteStream(
	ctx context.Context,
	request CompleteRequest,
) (stream *CompletionStream, err error) {

	request.Stream = true
	if len(request.Prompt) == 0 {
		err = ErrCompletePromptNotEmpty
		return
	}
	if request.MaxTokensToSample == 0 {
		err = ErrCompleteMaxTokensToSmapleNotEmpty
		return
	}
	urlSuffix := "/complete"

	req, err := c.newRequest(ctx, "POST", c.fullURL(urlSuffix, request.Model), withBody(request))
	if err != nil {
		return nil, err
	}

	resp, err := sendRequestStream[CompleteResponse](c, req)
	if err != nil {
		return
	}
	stream = &CompletionStream{
		streamReader: resp,
	}
	return
}
