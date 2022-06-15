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

	ch := make(chan *model.Post, 1)

	go postProcessor(ctx, ch, apiPosts)

	for c := range ch {
		response = append(response, c)
	}

	return response, nil
}

func postProcessor(ctx context.Context, ch chan *model.Post, data []v1.PostService) {

	for _, v := range data {

		v.BuildPostsData(ctx, model.Post{}, ch)
	}

	close(ch)
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {

	response := []*model.Comment{}

	service := v1.CommentService{}

	apiComments, err := service.GetAll(ctx)
	if err != nil {
		return response, err
	}

	for _, v := range apiComments {
		tmp := &model.Comment{
			ID:     v.ID,
			PostID: v.PostID,
			Name:   v.Name,
			Email:  v.Email,
			Body:   v.Body,
		}

		response = append(response, tmp)
	}

	return response, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {

	response := []*model.User{}

	service := v1.UserService{}

	apiUsers, err := service.GetAll(ctx)
	if err != nil {
		return response, err
	}

	for _, v := range apiUsers {
		tmp := model.User{
			ID:       v.ID,
			Name:     v.Name,
			Username: v.Username,
			Email:    v.Email,
			Address: &model.UserAddress{
				Street:  v.Address.Street,
				Suite:   v.Address.Suite,
				City:    v.Address.City,
				Zipcode: v.Address.Zipcode,
				Geo: &model.UserAddressGeo{
					Lat: v.Address.Geo.Lat,
					Lng: v.Address.Geo.Lng,
				},
			},
			Phone:   v.Phone,
			Website: v.Website,
			Company: &model.UserCompany{
				Name:        v.Company.Name,
				CatchPhrase: v.Company.CatchPhrase,
				Bs:          v.Company.Bs,
			},
		}

		response = append(response, &tmp)
	}

	return response, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {

	response := &model.User{}

	service := v1.UserService{}

	apiUser, err := service.GetByID(ctx, id)
	if err != nil {
		return response, err
	}

	response = &model.User{
		ID:       apiUser.ID,
		Name:     apiUser.Name,
		Username: apiUser.Username,
		Email:    apiUser.Email,
		Address: &model.UserAddress{
			Street:  apiUser.Address.Street,
			Suite:   apiUser.Address.Suite,
			City:    apiUser.Address.City,
			Zipcode: apiUser.Address.Zipcode,
			Geo: &model.UserAddressGeo{
				Lat: apiUser.Address.Geo.Lat,
				Lng: apiUser.Address.Geo.Lng,
			},
		},
		Phone:   apiUser.Phone,
		Website: apiUser.Website,
		Company: &model.UserCompany{
			Name:        apiUser.Company.Name,
			CatchPhrase: apiUser.Company.CatchPhrase,
			Bs:          apiUser.Company.Bs,
		},
	}

	return response, nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*model.Post, error) {

	response := &model.Post{}

	userService := v1.UserService{}

	service := v1.PostService{}

	apiPost, err := service.GetByID(ctx, id)
	if err != nil {
		return response, err
	}

	userData, err := userService.GetByID(ctx, apiPost.UserID)
	if err != nil {
		return response, err
	}

	quantityComments, err := apiPost.GetQuantityCommentsByPost(ctx)
	if err != nil {
		return response, err
	}

	response = &model.Post{
		ID:    apiPost.ID,
		Title: apiPost.Title,
		Body:  apiPost.Body,
		Author: &model.User{
			ID:       userData.ID,
			Name:     userData.Name,
			Username: userData.Username,
			Email:    userData.Email,
			Address: &model.UserAddress{
				Street:  userData.Address.Street,
				Suite:   userData.Address.Suite,
				City:    userData.Address.City,
				Zipcode: userData.Address.Zipcode,
				Geo: &model.UserAddressGeo{
					Lat: userData.Address.Geo.Lat,
					Lng: userData.Address.Geo.Lng,
				},
			},
			Phone:   userData.Phone,
			Website: userData.Website,
			Company: &model.UserCompany{
				Name:        userData.Company.Name,
				CatchPhrase: userData.Company.CatchPhrase,
				Bs:          userData.Company.Bs,
			},
		},
		CreatedAt:        v1.Randate().String(),
		QuantityComments: quantityComments,
	}

	return response, nil
}

func (r *queryResolver) Comment(ctx context.Context, id int) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPostsByUserID(ctx context.Context, id int) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCommentsByPostID(ctx context.Context, id int) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
