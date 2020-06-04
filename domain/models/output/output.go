package output

import "time"

// FavoriteGamesRes :
type FavoriteGamesRes struct {
	GameTitle string    `json:"game_title" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

// UserResBody :
type UserResBody struct {
	UserID        string             `json:"id" binding:"required"`
	UserName      string             `json:"user_name" binding:"required"`
	Icon          string             `json:"icon" binding:"required"`
	CreatedAt     time.Time          `json:"created_at" binding:"required"`
	FavoriteGames []FavoriteGamesRes `json:"favorite_games" binding:"required"`
}

// RoomResBody :
type RoomResBody struct {
	RoomID        string             `json:"id" binding:"required"`
	OwnerID       string             `json:"user_name" binding:"required"`
	Icon          string             `json:"icon" binding:"required"`
	Name          time.Time          `json:"created_at" binding:"required"`
	FavoriteGames []FavoriteGamesRes `json:"favorite_games" binding:"required"`
}
