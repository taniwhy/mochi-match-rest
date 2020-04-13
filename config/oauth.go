package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var conf *oauth2.Config

func Init() *oauth2.Config {
	c := Config
	conf = &oauth2.Config{
		ClientID:     c.GoogleOAuth.ClientID,
		ClientSecret: c.GoogleOAuth.ClientSecret,
		RedirectURL:  c.GoogleOAuth.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}

func GetLoginURL(state string) string {
	Init()
	return conf.AuthCodeURL(state)
}
