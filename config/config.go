package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}
}

// Config : todo
var Config config

// InitConf : todo
func InitConf() {
	// 設定ファイル名を記載
	viper.SetConfigName("config")
	// ファイルタイプ
	viper.SetConfigType("yml")
	// ファイルパスの設定。クロスプラットフォームで参照できるようにfilepathライブラリを使用
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "github.com", "taniwhy", "mochi-match-rest", "config"))
	// 環境変数から設定値を上書きできるように設定
	viper.AutomaticEnv()
	// conf読み取り
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file read error")
		fmt.Println(err)
		os.Exit(1)
	}
	// UnmarshalしてCにマッピング
	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}
