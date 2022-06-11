package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Maximo-Miranda/challenge-fullstack/go-graphql-service/graph/generated"
	"github.com/Maximo-Miranda/challenge-fullstack/go-graphql-service/graph/model"
	dapr "github.com/dapr/go-sdk/client"
)

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {

	response := []*model.Post{}

	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	resp, err := client.InvokeMethod(ctx, "jsonplaceholder", "posts/1", "get")
	if err != nil {
		return response, err
	}

	fmt.Println(string(resp))
	
	response = append(response, &model.Post{
		ID:    1,
		Title: "Title test",
		Body:  "Body test",
		Author: &model.User{
			ID:       0,
			Name:     "Maximo Miranda",
			Username: "mmiranda",
			Email:    "maximo@gmail.com",
			Address:  nil,
			Phone:    "300123456",
			Website:  "http://localhost",
			Company:  nil,
		},
		CreatedAt:        "2022-06-10",
		QuantityComments: 15,
	})

	return response, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
