package input

// UserCreateReqBody :
type UserCreateReqBody struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// FavoriteGameRecord :
type FavoriteGameRecord struct {
	GameTitle string `json:"game_title" binding:"required"`
}

// UserUpdateReqBody :
type UserUpdateReqBody struct {
	UserName      string               `json:"user_name" binding:"required"`
	Icon          string               `json:"icon" binding:"required"`
	FavoriteGames []FavoriteGameRecord `json:"favorite_games" binding:"required"`
}

// RoomCreateReqBody :
type RoomCreateReqBody struct {
	RoomText    string `json:"room_text" binding:"required"`
	GameTitleID string `json:"game_title_id" binding:"required"`
	GameHardID  string `json:"game_hard_id" binding:"required"`
	Capacity    int    `json:"capacity" binding:"required"`
}
