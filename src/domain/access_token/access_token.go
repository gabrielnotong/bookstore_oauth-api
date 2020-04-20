package access_token

import (
	"github.com/gabrielnotong/bookstore_oauth-api/src/errors"
	"strings"
	"time"
)

const (
	expirationDate = 24
)

type AccessToken struct {
	ID       string `json:"id"`
	UserId   int64  `json:"user_id"`
	ClientId int64  `json:"client_id"` // to define token duration
	Expires  int64  `json:"expires"`
}

func NewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationDate * time.Hour).Unix(),
	}
}

func (at AccessToken) Validate() *errors.RestErr {
	if strings.TrimSpace(at.ID) == "" {
		return errors.NewBadRequestError("invalid token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration date")
	}

	return nil
}

func (at *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationDate := time.Unix(at.Expires, 0)
	return now.After(expirationDate)
}
