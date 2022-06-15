package v1

import (
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
)

type CommentService struct {
	ID     int    `json:"id"`
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func (m *CommentService) GetAll(ctx context.Context) ([]CommentService, error) {

	response := []CommentService{}

	client, err := dapr.NewClient()
	if err != nil {
		return response, err
	}

	resp, err := client.InvokeMethod(ctx, "jsonplaceholder", "comments", "get")
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(resp, &response); err != nil {
		return response, err
	}

	return response, nil
}

func (m *CommentService) GetCommentsByPostID(ctx context.Context, postID int) ([]CommentService, error) {

	response := []CommentService{}

	client, err := dapr.NewClient()
	if err != nil {
		return response, err
	}

	resp, err := client.InvokeMethod(ctx, "jsonplaceholder", fmt.Sprintf("/comments?postId=%d", postID), "get")
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(resp, &response); err != nil {
		return response, err
	}

	return response, nil
}
