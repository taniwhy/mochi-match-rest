package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	switch os.Getenv("GO_ENV") {
	case "dev":
		err := godotenv.Load(
			fmt.Sprintf("%s/src/github.com/taniwhy/mochi-match-rest/config/env/.env.dev", os.Getenv("GOPATH")))
		if err != nil {
			panic(err)
		}
	case "test":
		err := godotenv.Load(
			fmt.Sprint("/drone/src/config/env/.env.test"))
		if err != nil {
			panic(err)
		}
	case "prod":
		err := godotenv.Load(
			fmt.Sprintf("%s/src/github.com/taniwhy/mochi-match-rest/config/env/.env", os.Getenv("GOPATH")))
		if err != nil {
			panic(err)
		}
	default:
		err := godotenv.Load(
			fmt.Sprintf("%s/src/github.com/taniwhy/mochi-match-rest/config/env/.env.dev", os.Getenv("GOPATH")))
		if err != nil {
			panic(err)
		}
	}

}

// GetDatabaseConf :　データベースの接続情報の取得
func GetDatabaseConf() (dsn string) { return os.Getenv("DB_URL") }

// GetRedisConf : Reidsの接続情報の取得
func GetRedisConf() (size int, network, addr, pass, key string) {
	size, _ = strconv.Atoi(os.Getenv("REDIS_SIZE"))
	network = os.Getenv("REDIS_NETWORK")
	addr = os.Getenv("REDIS_ADDR")
	pass = os.Getenv("REDIS_PASS")
	key = os.Getenv("REDIS_KEY")
	return
}

// GetOAuthClientConf : GoogleAuth認証情報の取得
func GetOAuthClientConf() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
