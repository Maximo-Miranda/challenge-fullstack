// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID     int    `json:"id"`
	PostID int    `json:"postId"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

type Post struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Body             string `json:"body"`
	Author           *User  `json:"author"`
	CreatedAt        string `json:"createdAt"`
	QuantityComments int    `json:"quantityComments"`
}

type User struct {
	ID       int          `json:"id"`
	Name     string       `json:"name"`
	Username string       `json:"username"`
	Email    string       `json:"email"`
	Address  *UserAddress `json:"address"`
	Phone    string       `json:"phone"`
	Website  string       `json:"website"`
	Company  *UserCompany `json:"company"`
}

type UserAddress struct {
	Street  string          `json:"street"`
	Suite   string          `json:"suite"`
	City    string          `json:"city"`
	Zipcode string          `json:"zipcode"`
	Geo     *UserAddressGeo `json:"geo"`
}

type UserAddressGeo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type UserCompany struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}
