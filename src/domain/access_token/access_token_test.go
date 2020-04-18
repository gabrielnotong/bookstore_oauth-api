package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestNewAccessToken(t *testing.T) {
	at := NewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, "", at.Token, "new access token should not have defined token id")
	assert.True(t, at.UserId == 0, "new access token should not have an associated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := &AccessToken{}
	assert.True(t, at.IsExpired(), "empty token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should NOT be expired")
}
