package access_token

import "time"

const (
	expirationTime = 24
)

type AccessToken struct {
	Token    string `json:"token"`
	UserId   int64  `json:"user_id"`
	ClientId int64  `json:"client_id"` // to define token duration
	Expires  int64  `json:"expires"`
}

func NewAccessToken() *AccessToken  {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) IsExpired() bool {
	return time.Now().UTC().After(time.Unix(at.Expires, 0))
}