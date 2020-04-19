package db

import (
	"github.com/gabrielnotong/bookstore_oauth-api/src/domain/access_token"
	"github.com/gabrielnotong/bookstore_oauth-api/src/errors"
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
	// define cassandra db entry point here
}

func NewDBRepository() DBRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetById(tokenId string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("Database access not yet implemented")
}
