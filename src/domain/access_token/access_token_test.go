package access_token

import (
	"testing"
	"time"
)

func TestExpirationDate(t *testing.T) {
	if expirationDate != 24 {
		t.Error("Expiration date should be 24")
	}
}

func TestNewAccessToken(t *testing.T) {
	at := NewAccessToken()
	if at.IsExpired() {
		t.Error("New access token should not be expired")
	}

	if at.ID != "" {
		t.Error("New access token should not have ID")
	}

	if at.UserId != 0 {
		t.Error("New access token should not have user id attached to it")
	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := &AccessToken{}

	if !at.IsExpired() {
		t.Error("Access token should be expired")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("Access token should expire after 24 hours")
	}
}
