package db

import (
	"github.com/gabrielnotong/bookstore_oauth-api/src/domain/access_token"
	"github.com/gabrielnotong/bookstore_oauth-api/src/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT id, user_id, client_id, expires FROM access_token WHERE id = ?"
	queryCreateAccessToken = "INSERT INTO access_token(id, user_id, client_id, expires) VALUES (?,?,?,?)"
)

type DBRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
	session *gocql.Session
}

func NewDBRepository(s *gocql.Session) DBRepository {
	return &dbRepository{s}
}

func (db *dbRepository) GetById(tokenId string) (*access_token.AccessToken, *errors.RestErr) {
	q := db.session.Query(queryGetAccessToken, tokenId)
	at := &access_token.AccessToken{}
	err := q.Scan(&at.ID, &at.UserId, &at.ClientId, &at.Expires)
	if err != nil {
		return nil, errors.ParsePostgresError(err)
	}
	return at, nil
}

func (db *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	err := db.session.Query(queryCreateAccessToken, at.ID, at.UserId, at.ClientId, at.Expires).Exec()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
