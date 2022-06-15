package v1

import (
	"context"
	"encoding/json"
	dapr "github.com/dapr/go-sdk/client"
)

type PostService struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (m *PostService) GetAll(ctx context.Context) ([]PostService, error) {

	response := []PostService{}

	client, err := dapr.NewClient()
	if err != nil {
		return response, err
	}

	resp, err := client.InvokeMethod(ctx, "jsonplaceholder", "posts", "get")
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(resp, &response); err != nil {
		return response, err
	}

	return response, nil
}
