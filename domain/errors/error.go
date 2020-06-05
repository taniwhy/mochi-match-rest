package errors

import (
	"fmt"
	"strings"
	"time"

	"github.com/taniwhy/mochi-match-rest/domain/models/input"
)

// ErrUserCreateReqBinding :
type ErrUserCreateReqBinding struct {
	UserName string
	Email    string
}

func (b ErrUserCreateReqBinding) Error() string {
	var errMsg []string
	if b.UserName == "" {
		errMsg = append(errMsg, "user_name")
	}
	if b.Email == "" {
		errMsg = append(errMsg, "email")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrUserUpdateReqBinding :
type ErrUserUpdateReqBinding struct {
	UserName      string
	Icon          string
	FavoriteGames []input.FavoriteGameRecord
}

func (b ErrUserUpdateReqBinding) Error() string {
	var errMsg []string
	if b.UserName == "" {
		errMsg = append(errMsg, "user_name")
	}
	if b.Icon == "" {
		errMsg = append(errMsg, "icon")
	}
	if b.FavoriteGames == nil {
		errMsg = append(errMsg, "favorite_games")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrRoomCreateReqBinding :
type ErrRoomCreateReqBinding struct {
	RoomText    string
	GameTitleID string
	GameHardID  string
	Capacity    int
	Start       time.Time
}

func (b ErrRoomCreateReqBinding) Error() string {
	var errMsg []string
	if b.RoomText == "" {
		errMsg = append(errMsg, "room_text")
	}
	if b.GameTitleID == "" {
		errMsg = append(errMsg, "game_title_id")
	}
	if b.GameHardID == "" {
		errMsg = append(errMsg, "game_hard_id")
	}
	if b.Capacity == 0 {
		errMsg = append(errMsg, "capacity")
	}
	if b.Start.IsZero() == true {
		errMsg = append(errMsg, "start")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrGenerateID :
type ErrGenerateID struct{}

func (b ErrGenerateID) Error() string {
	return fmt.Sprintf("Failed generate ID")
}

// ErrCoockie :
type ErrCoockie struct {
	Key   string
	Value string
}

func (c ErrCoockie) Error() string {
	return fmt.Sprintf("Coockie error! - target: %v - got: %v", c.Key, c.Value)
}

// ErrParams :
type ErrParams struct {
	Need string
	Got  string
}

func (c ErrParams) Error() string {
	return fmt.Sprintf("Params error! - need: %v - got: %v", c.Need, c.Got)
}

// ErrGetTokenClaims :
type ErrGetTokenClaims struct {
	Detail interface{}
}

func (c ErrGetTokenClaims) Error() string {
	return fmt.Sprintf("Failed token claims! - details: %v", c.Detail)
}

// ErrIDAlreadyExists :
type ErrIDAlreadyExists struct {
	Provider string
	ID       string
}

func (e ErrIDAlreadyExists) Error() string {
	return fmt.Sprintf("ID already exists! - provider: %s - ID: %s", e.Provider, e.ID)
}

// ErrRoomAlreadyExists :
type ErrRoomAlreadyExists struct{}

func (e ErrRoomAlreadyExists) Error() string {
	return fmt.Sprintf("Room already exists!")
}

// ErrDataBase :
type ErrDataBase struct {
	Detail interface{}
}

func (e ErrDataBase) Error() string {
	return fmt.Sprintf("Database error! - detail: %s", e.Detail)
}

// ErrNotFound :
type ErrNotFound struct {
	Detail interface{}
}

func (n ErrNotFound) Error() string {
	return fmt.Sprintf("Not found!")
}

// ErrRecordNotFound :
type ErrRecordNotFound struct {
	Detail interface{}
}

func (e ErrRecordNotFound) Error() string {
	return fmt.Sprintf("Record not found! - detail: %s", e.Detail)
}

// ErrUnexpectedQueryProvider :
type ErrUnexpectedQueryProvider struct {
	Provider string
}

func (p ErrUnexpectedQueryProvider) Error() string {
	return fmt.Sprintf("Unexpected query provider! - provider: %s", p.Provider)
}
