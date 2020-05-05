package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		URL:          "https://api.bookstore.com/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      "{'email':'email@yopmail.com','password':'password'}",
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := NewUsersRepository()

	user, err := repository.Login("email@yopmail.com", "password")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid rest client response when trying to log user in", err.Message)
}
