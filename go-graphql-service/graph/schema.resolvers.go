package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Maximo-Miranda/challenge-fullstack/go-graphql-service/graph/generated"
	"github.com/Maximo-Miranda/challenge-fullstack/go-graphql-service/graph/model"
	"github.com/Maximo-Miranda/challenge-fullstack/go-graphql-service/internal/services/v1"
)

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {

	response := []*model.Post{}

	service := v1.PostService{}

	apiPosts, err := service.GetAll(ctx)
	if err != nil {
		return response, err
	}

	for _, v := range apiPosts {
		tmp := model.Post{
			ID:               int(v.ID),
			Title:            v.Title,
			Body:             v.Body,
			Author:           nil,
			CreatedAt:        "2022-06-10",
			QuantityComments: 15,
		}

		response = append(response, &tmp)
	}

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
