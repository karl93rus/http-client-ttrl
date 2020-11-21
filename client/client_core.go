package client

import (
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	fullHeaders := c.getClientHeaders(headers)
	request.Header = fullHeaders

	return client.Do(request)
}

func (c *httpClient) getClientHeaders(headers http.Header) http.Header {
	result := make(http.Header)

	// Set common headers
	for h, v := range c.Headers {
		if len(v) > 0 {
			result.Set(h, v[0])
		}
	}

	// Set method headers
	for h, v := range headers {
		if len(v) > 0 {
			result.Set(h, v[0])
		}
	}

	return result
}
