package errors

import (
	"fmt"

	"github.com/taniwhy/mochi-match-rest/domain/models/input"
)

// ErrGenerateID :
type ErrGenerateID struct{}

func (b ErrGenerateID) Error() string {
	return fmt.Sprintf("Failed generate ID")
}

// ErrCreateReqBinding :
type ErrCreateReqBinding struct {
	UserName string
	Email    string
}

func (b ErrCreateReqBinding) Error() string {
	if b.UserName == "" && b.Email == "" {
		return fmt.Sprintf("Binding error! - user_name, email is required")
	} else if b.Email == "" {
		return fmt.Sprintf("Binding error! - email is required")
	} else {
		return fmt.Sprintf("Binding error! - user_name is required")
	}
}

// ErrUpdateReqBinding :
type ErrUpdateReqBinding struct {
	UserName      string
	Icon          string
	FavoriteGames []input.FavoriteGameRecord
}

func (b ErrUpdateReqBinding) Error() string {
	if b.UserName == "" && b.Icon == "" && b.FavoriteGames == nil {
		return fmt.Sprintf("Binding error! - user_name, icon, favorite_games is required")
	} else if b.UserName == "" && b.Icon == "" {
		return fmt.Sprintf("Binding error! - user_name, icon is required")
	} else if b.UserName == "" && b.FavoriteGames == nil {
		return fmt.Sprintf("Binding error! - user_name, favorite_games is required")
	} else if b.Icon == "" && b.FavoriteGames == nil {
		return fmt.Sprintf("Binding error! - icon, favorite_games is required")
	} else if b.UserName == "" {
		return fmt.Sprintf("Binding error! - user_name is required")
	} else if b.Icon == "" {
		return fmt.Sprintf("Binding error! - icon is required")
	} else {
		return fmt.Sprintf("Binding error! - favorite_games is required")
	}
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

// ErrDataBase :
type ErrDataBase struct {
	Detail interface{}
}

func (e ErrDataBase) Error() string {
	return fmt.Sprintf("Database error! - detail: %s", e.Detail)
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
