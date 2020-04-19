package access_token

import (
	"github.com/gabrielnotong/bookstore_oauth-api/src/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetById(tokenId string) (*AccessToken, *errors.RestErr) {

	if strings.TrimSpace(tokenId) == "" {
		return nil, errors.NewBadRequestError("Invalid token id")
	}

	return s.r.GetById(tokenId)
}
