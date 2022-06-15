package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Maximo-Miranda/challenge-fullstack/go-graphql-service/graph/model"
	dapr "github.com/dapr/go-sdk/client"
	"math/rand"
	"time"
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

func (m *PostService) GetQuantityCommentsByPost(ctx context.Context) (int, error) {

	quantity := 0

	commentService := CommentService{}

	resp, err := commentService.GetCommentsByPostID(ctx, m.ID)
	if err != nil {
		return quantity, err
	}

	quantity = len(resp)

	return quantity, nil
}

func (m *PostService) GetByID(ctx context.Context, id int) (PostService, error) {

	response := PostService{}

	client, err := dapr.NewClient()
	if err != nil {
		return response, err
	}

	resp, err := client.InvokeMethod(ctx, "jsonplaceholder", fmt.Sprintf("posts/%d", id), "get")
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(resp, &response); err != nil {
		return response, err
	}

	return response, nil

}

func (m *PostService) BuildPostsData(ctx context.Context, post model.Post, ch chan<- *model.Post) {

	userService := UserService{}

	userData, err := userService.GetByID(ctx, m.UserID)
	if err != nil {
		ch <- nil
	}

	quantityComments, err := m.GetQuantityCommentsByPost(ctx)
	if err != nil {
		ch <- nil
	}

	post.ID = m.ID
	post.Body = m.Body
	post.Title = m.Title
	post.Author = &model.User{
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
	}
	post.CreatedAt = Randate().String()
	post.QuantityComments = quantityComments

	ch <- &post
}

func Randate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
