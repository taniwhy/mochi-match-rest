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
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

// GetDatabaseConf :
func GetDatabaseConf() (dsn string) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	dsn = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname,
	)
	return
}

// GetRedisConf :
func GetRedisConf() (size int, network, addr, pass, key string) {
	size, _ = strconv.Atoi(os.Getenv("REDIS_SIZE"))
	network = os.Getenv("REDIS_NETWORK")
	addr = os.Getenv("REDIS_ADDR")
	pass = os.Getenv("REDIS_PASS")
	key = os.Getenv("REDIS_KEY")
	return
}

// ConfigureOAuthClient :
func ConfigureOAuthClient() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_REDIRECT_URL"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		RedirectURL:  os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
