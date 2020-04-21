package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type config struct {
	Database struct {
		Host   string
		Port   string
		User   string
		Pass   string
		DBName string
	}
	Redis struct {
		Size    int
		Network string
		Addr    string
		Pass    string
		Key     string
	}
	GoogleOAuth struct {
		RedirectURL  string
		ClientID     string
		ClientSecret string
	}
}

// config : todo
var (
	Config config
)

// InitConf : todo
func InitConf() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	// ファイルパスの設定。クロスプラットフォームで参照できるようにfilepathライブラリを使用
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "github.com", "taniwhy", "mochi-match-rest", "config"))
	// 環境変数から設定値を上書きできるように設定
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file read error")
		fmt.Println(err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}

// GetDatabaseConf :
func GetDatabaseConf() (dsn string) {
	host := Config.Database.Host
	port := Config.Database.Port
	user := Config.Database.User
	pass := Config.Database.Pass
	dbname := Config.Database.DBName
	dsn = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname,
	)
	return
}

// GetRedisConf :
func GetRedisConf() (size int, network, addr, pass, key string) {
	size = Config.Redis.Size
	network = Config.Redis.Network
	addr = Config.Redis.Addr
	pass = Config.Redis.Pass
	key = Config.Redis.Key
	return
}

// ConfigureOAuthClient :
func ConfigureOAuthClient() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     Config.GoogleOAuth.ClientID,
		ClientSecret: Config.GoogleOAuth.ClientSecret,
		RedirectURL:  Config.GoogleOAuth.RedirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
