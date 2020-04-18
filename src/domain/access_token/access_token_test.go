package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstant(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}
}

func TestNewAccessToken(t *testing.T) {
	at := NewAccessToken()
	if at.IsExpired() {
		t.Error("brand new access token should not be expired")
	}

	if at.Token != "" {
		t.Error("new access token should not have defined token id")
	}

	if at.UserId != 0 {
		t.Error("new access token should not have an associated user id")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := &AccessToken{}

	if !at.IsExpired() {
		t.Error("empty token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token expiring three hours from now should NOT be expired")
	}
}
