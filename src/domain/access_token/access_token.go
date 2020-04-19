package access_token

import "time"

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

func (at *AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationDate := time.Unix(at.Expires, 0)
	return now.After(expirationDate)
}
