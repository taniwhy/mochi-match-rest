package input

import "gopkg.in/guregu/null.v4"

// UserCreateBody : ユーザー作成リクエストボディ
type UserCreateBody struct {
	Provider   string
	ProviderID string
	UserName   string
	Email      string
}

// FavoriteGameRecord : お気に入りゲームレコード
type FavoriteGameRecord struct {
	GameTitle string `json:"game_title" binding:"required"`
}

// UserUpdateReqBody : ユーザー更新リクエストボディ
type UserUpdateReqBody struct {
	UserName      string               `json:"user_name" binding:"required"`
	Icon          string               `json:"icon" binding:"required"`
	FavoriteGames []FavoriteGameRecord `json:"favorite_games" binding:"required"`
}

// RoomCreateReqBody :ルーム作成リクエストボディ
type RoomCreateReqBody struct {
	RoomText    string    `json:"room_text" binding:"required"`
	GameTitleID string    `json:"game_title_id" binding:"required"`
	GameHardID  string    `json:"game_hard_id" binding:"required"`
	Capacity    int       `json:"capacity" binding:"required"`
	Start       null.Time `json:"start" binding:"required"`
}

// TokenReqBody : トークンのリクエストボディ
type TokenReqBody struct {
	RefreshToken string `json:"refresh_token"`
}
