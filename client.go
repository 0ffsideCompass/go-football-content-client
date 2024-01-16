package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Client struct {
	domain string
	client *http.Client
}

// NewClients creates a client object and check if the information provided are not empty
func NewClient(domain, apiKey string) (*Client, error) {
	switch {
	case domain == ``:
		return nil, fmt.Errorf(`error_no_domain`)
	}

	return &Client{
		domain: domain,
		client: &http.Client{},
	}, nil
}

func (c *Client) post(endpoint string, payload string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.domain, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	return body, err
}
