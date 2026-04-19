package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// HTTPClient data model
type HTTPClient struct {
	httpClient *http.Client
}

// NewRequest function for intialize httpRequest object
// Paramter, timeout in time.Duration
func NewRequest(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		httpClient: &http.Client{Timeout: time.Second * timeout},
	}
}

// Call private function for call http request
func (c *HTTPClient) Call(method, fullURL string, body io.Reader, response any, headers map[string]string) error {
	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(response)
}
