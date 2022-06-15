package v1

import (
	"context"
	"encoding/json"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
)

type addressGeo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type userAddress struct {
	Street  string     `json:"street"`
	Suite   string     `json:"suite"`
	City    string     `json:"city"`
	Zipcode string     `json:"zipcode"`
	Geo     addressGeo `json:"geo"`
}

type userCompany struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type UserService struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Address  userAddress `json:"address"`
	Phone    string      `json:"phone"`
	Website  string      `json:"website"`
	Company  userCompany `json:"company"`
}

// GetAll ...
func (m *UserService) GetAll(ctx context.Context) ([]UserService, error) {

	response := []UserService{}

	client, err := dapr.NewClient()
	if err != nil {
		return response, err
	}

	resp, err := client.InvokeMethod(ctx, "jsonplaceholder", "users", "get")
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(resp, &response); err != nil {
		return response, err
	}

	return response, nil

}

// GetByID ...
func (m *UserService) GetByID(ctx context.Context, id int) (UserService, error) {

	response := UserService{}

	client, err := dapr.NewClient()
	if err != nil {
		return response, err
	}

	resp, err := client.InvokeMethod(ctx, "jsonplaceholder", fmt.Sprintf("users/%d", id), "get")
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(resp, &response); err != nil {
		return response, err
	}

	return response, nil

}
