package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExpirationDate(t *testing.T) {
	assert.True(t, expirationDate == 24, "Expiration date should be 24")
}

func TestNewAccessToken(t *testing.T) {
	at := NewAccessToken()
	assert.False(t, at.IsExpired(), "New access token should not be expired")
	assert.EqualValues(t, "", at.ID, "New access token should not have ID")
	assert.False(t, at.UserId != 0, "New access token should not have user id attached to it")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := &AccessToken{}
	assert.True(t, at.IsExpired(), "Access token should be expired")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Access token should expire after 24 hours")
}
