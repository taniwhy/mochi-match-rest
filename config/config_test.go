package config

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2/google"
)

func TestGetDatabaseConf(t *testing.T) {
	assert.NotEqual(t, GetDatabaseConf(), "")
	assert.Equal(t, GetDatabaseConf(), os.Getenv("DB_URL"))
}

func TestGetRedisConf(t *testing.T) {
	size, network, addr, pass, key := GetRedisConf()
	sizeEnv, _ := strconv.Atoi(os.Getenv("REDIS_SIZE"))

	assert.NotEqual(t, size, 0)
	assert.Equal(t, size, sizeEnv)

	assert.NotEqual(t, network, "")
	assert.Equal(t, network, os.Getenv("REDIS_NETWORK"))

	assert.NotEqual(t, addr, "")
	assert.Equal(t, addr, os.Getenv("REDIS_ADDR"))

	assert.Equal(t, pass, os.Getenv("REDIS_PASS"))

	assert.NotEqual(t, key, "")
	assert.Equal(t, key, os.Getenv("REDIS_KEY"))
}

func TestGetOAuthClientConf(t *testing.T) {
	conf := GetOAuthClientConf()

	assert.NotEqual(t, conf.ClientID, "")
	assert.Equal(t, conf.ClientID, os.Getenv("GOOGLE_OAUTH_CLIENT_ID"))

	assert.NotEqual(t, conf.ClientSecret, "")
	assert.Equal(t, conf.ClientSecret, os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"))

	assert.NotEqual(t, conf.RedirectURL, "")
	assert.Equal(t, conf.RedirectURL, os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"))

	assert.NotEqual(t, conf.Scopes[0], "")
	assert.NotEqual(t, conf.Scopes[1], "")
	assert.Equal(t, conf.Scopes[0], "email")
	assert.Equal(t, conf.Scopes[1], "profile")
	assert.Equal(t, len(conf.Scopes), 2)

	assert.NotEqual(t, conf.Endpoint, "")
	assert.Equal(t, conf.Endpoint, google.Endpoint)
}
