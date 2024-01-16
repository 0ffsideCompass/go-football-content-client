package client

import (
	"encoding/json"
	"fmt"
)

const (
	fixturesByDate  = "/fixtures/date"
	generateArticle = "/article/generate"
)

func (c *Client) FixtureByDate(date string) (*FixtureByDateResponse, error) {
	var resp FixtureByDateResponse
	req := FixtureByDateRequest{Date: date}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	body, err := c.post(fixturesByDate, string(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &resp, nil
}

func (c *Client) GenerateArticle(id int, contentType, language string) (*ArticleResponse, error) {
	var resp ArticleResponse
	req := ArticleGenerateRequest{
		ID:          id,
		ContentType: contentType,
		Language:    language,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	body, err := c.post(generateArticle, string(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error posting request: %w", err)
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return &resp, nil
}
