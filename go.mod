module github.com/taniwhy/mochi-match-rest

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/go-chi/chi v4.1.2+incompatible // indirect
	github.com/golang/mock v1.4.3
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/jinzhu/gorm v1.9.12
	github.com/k-washi/gologger v0.0.0-20200119183124-3c165c038726 // indirect
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.7.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
)

replace github.com/taniwhy/mochi-match-rest => ./
