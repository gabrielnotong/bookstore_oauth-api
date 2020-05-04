package rest

import (
	"encoding/json"
	"github.com/gabrielnotong/bookstore_oauth-api/src/domain/users"
	"github.com/gabrielnotong/bookstore_oauth-api/src/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var usersRestClient = rest.RequestBuilder{
	BaseURL:        "https://api.bookstore.com",
	Timeout:        100 * time.Millisecond,
}

type RestUsersRepository interface {
	Login(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (ur *usersRepository) Login(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	response := usersRestClient.Post("/users/login", request)
	// this checks for response timeout
	if response == nil || response.Response == nil {
		return nil, errors.NewBadRequestError("Invalid rest client response when trying to log user in")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface when trying to log in user")
		}
		return nil, &restErr
	}

	var u users.User
	if err := json.Unmarshal(response.Bytes(), &u); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal the user response")
	}

	return &u, nil
}
