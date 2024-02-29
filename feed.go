package client

import (
	"encoding/json"
	"fmt"
)

const (
	feedEndpoint = "/feed"
)

func (c *Client) Feed(leagues []int) (*FeedData, error) {
	var resp FeedData
	req := FeedRequest{Leagues: leagues}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	body, err := c.post(feedEndpoint, string(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &resp, nil
}
