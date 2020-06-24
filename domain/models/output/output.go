package output

import "time"

// FavoriteGamesRes : お気に入りゲームレスポンス
type FavoriteGamesRes struct {
	GameTitle string    `json:"game_title" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

// UserResBody : ユーザーレスポンス
type UserResBody struct {
	UserID        string             `json:"id" binding:"required"`
	UserName      string             `json:"user_name" binding:"required"`
	Icon          string             `json:"icon" binding:"required"`
	CreatedAt     time.Time          `json:"created_at" binding:"required"`
	FavoriteGames []FavoriteGamesRes `json:"favorite_games" binding:"required"`
}

// RoomResBody : ルームレスポンス
type RoomResBody struct {
	RoomID        string             `json:"id" binding:"required"`
	OwnerID       string             `json:"user_name" binding:"required"`
	Icon          string             `json:"icon" binding:"required"`
	Name          time.Time          `json:"created_at" binding:"required"`
	FavoriteGames []FavoriteGamesRes `json:"favorite_games" binding:"required"`
}
