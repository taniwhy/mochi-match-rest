package input

// CreateReqBody :
type CreateReqBody struct {
	UserName string `json:"user_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// FavoriteGameRecord :
type FavoriteGameRecord struct {
	GameTitle string `json:"game_title" binding:"required"`
}

// UpdateReqBody :
type UpdateReqBody struct {
	UserName      string               `json:"user_name" binding:"required"`
	Icon          string               `json:"icon" binding:"required"`
	FavoriteGames []FavoriteGameRecord `json:"favorite_games" binding:"required"`
}
