package errors

import (
	"fmt"
	"strings"
	"time"

	"github.com/taniwhy/mochi-match-rest/domain/models/input"
)

// ErrUserCreateReqBinding : ユーザー作成リクエストボディのバインディングエラー
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

// ErrUserUpdateReqBinding : ユーザー更新リクエストボディのバインディングエラー
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

// ErrRoomCreateReqBinding : ルーム作成リクエストボディのバインディングエラー
type ErrRoomCreateReqBinding struct {
	RoomText   string
	GameListID string
	GameHardID string
	Capacity   int
	Start      time.Time
}

func (b ErrRoomCreateReqBinding) Error() string {
	var errMsg []string
	if b.RoomText == "" {
		errMsg = append(errMsg, "room_text")
	}
	if b.GameListID == "" {
		errMsg = append(errMsg, "game_list_id")
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

// ErrGameListCreateReqBinding : ゲームリスト作成リクエストボディのバインディングエラー
type ErrGameListCreateReqBinding struct {
	GameTitle string
}

func (b ErrGameListCreateReqBinding) Error() string {
	var errMsg []string
	if b.GameTitle == "" {
		errMsg = append(errMsg, "game_title")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrGameListUpdateReqBinding : ゲームリスト更新リクエストボディのバインディングエラー
type ErrGameListUpdateReqBinding struct {
	GameTitle string
}

func (b ErrGameListUpdateReqBinding) Error() string {
	var errMsg []string
	if b.GameTitle == "" {
		errMsg = append(errMsg, "game_title")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrGameHardCreateReqBinding : ゲームハード作成リクエストボディのバインディングエラー
type ErrGameHardCreateReqBinding struct {
	HardName string
}

func (b ErrGameHardCreateReqBinding) Error() string {
	var errMsg []string
	if b.HardName == "" {
		errMsg = append(errMsg, "hard_name")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrGameHardUpdateReqBinding : ゲームハード更新リクエストボディのバインディングエラー
type ErrGameHardUpdateReqBinding struct {
	HardName string
}

func (b ErrGameHardUpdateReqBinding) Error() string {
	var errMsg []string
	if b.HardName == "" {
		errMsg = append(errMsg, "hard_name")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrReportReqBinding : レポートリクエストボディのバインディングエラー
type ErrReportReqBinding struct {
	VaiolatorID      string
	VaiolationDetail string
}

func (b ErrReportReqBinding) Error() string {
	var errMsg []string
	if b.VaiolatorID == "" {
		errMsg = append(errMsg, "vaiolator_id")
	}
	if b.VaiolationDetail == "" {
		errMsg = append(errMsg, "detail")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrChatPostCreateReqBinding : チャットメッセージ作成リクエストボディのバインディングエラー
type ErrChatPostCreateReqBinding struct {
	Message string
}

func (b ErrChatPostCreateReqBinding) Error() string {
	var errMsg []string
	if b.Message == "" {
		errMsg = append(errMsg, "message")
	}
	errMsgs := strings.Join(errMsg, ", ")
	return fmt.Sprintf("Binding error! - " + errMsgs + " is required")
}

// ErrSessionSave : セッションのセーブエラー
type ErrSessionSave struct{}

func (b ErrSessionSave) Error() string {
	return fmt.Sprintf("Failed session save")
}

// ErrInvalidSessionState : 無効のセッションステートエラー
type ErrInvalidSessionState struct {
	State interface{}
}

func (b ErrInvalidSessionState) Error() string {
	return fmt.Sprintf("Invalid session state: %s", b.State)
}

// ErrGoogleOAuthTokenExchange : GoogleOAuthトークンの変換エラー
type ErrGoogleOAuthTokenExchange struct{}

func (b ErrGoogleOAuthTokenExchange) Error() string {
	return fmt.Sprintf("Failed GoogleOAuth token exchange")
}

// ErrInvalidGoogleOAuthToken : 無効のGoogleOAuthトークンエラー
type ErrInvalidGoogleOAuthToken struct{}

func (b ErrInvalidGoogleOAuthToken) Error() string {
	return fmt.Sprintf("Invalid GoogleOAuth token")
}

// ErrGoogleAPIRequest : GoogleAPIのリクエストエラー
type ErrGoogleAPIRequest struct{}

func (b ErrGoogleAPIRequest) Error() string {
	return fmt.Sprintf("Failed GoogleAPI request")
}

// ErrReadGoogleAPIResponse : GoogleAPIのレスポンス読み込みエラー
type ErrReadGoogleAPIResponse struct{}

func (b ErrReadGoogleAPIResponse) Error() string {
	return fmt.Sprintf("Failed read GoogleAPI response")
}

// ErrUnmarshalJSON : ErrUnmarshalJSONエラー
type ErrUnmarshalJSON struct{}

func (b ErrUnmarshalJSON) Error() string {
	return fmt.Sprintf("Failed unmarshal json")
}

// ErrGenerateID : IDの生成エラー
type ErrGenerateID struct{}

func (b ErrGenerateID) Error() string {
	return fmt.Sprintf("Failed generate ID")
}

// ErrCoockie : クッキーエラー
type ErrCoockie struct {
	Key   string
	Value string
}

func (c ErrCoockie) Error() string {
	return fmt.Sprintf("Coockie error! - target: %v - got: %v", c.Key, c.Value)
}

// ErrParams : HTTPリクエストパラメータエラー
type ErrParams struct {
	Need string
	Got  string
}

func (c ErrParams) Error() string {
	return fmt.Sprintf("Params error! - need: %v - got: %v", c.Need, c.Got)
}

// ErrGetTokenClaims : トークンClaimsの取得エラー
type ErrGetTokenClaims struct {
	Detail interface{}
}

func (c ErrGetTokenClaims) Error() string {
	return fmt.Sprintf("Failed token claims! - details: %v", c.Detail)
}

// ErrIDAlreadyExists : IDの重複エラー
type ErrIDAlreadyExists struct {
	Provider string
	ID       string
}

func (e ErrIDAlreadyExists) Error() string {
	return fmt.Sprintf("ID already exists! - provider: %s - ID: %s", e.Provider, e.ID)
}

// ErrRoomAlreadyExists : ルームの重複エラー
type ErrRoomAlreadyExists struct{}

func (e ErrRoomAlreadyExists) Error() string {
	return fmt.Sprintf("Room already exists!")
}

// ErrRoomAlreadyLock : ロック済みルームへの入室リクエストエラー
type ErrRoomAlreadyLock struct {
	RoomID string
}

func (l ErrRoomAlreadyLock) Error() string {
	return fmt.Sprintf("Room already lock! - room: %s", l.RoomID)
}

// ErrRoomCapacityOver : ルームの許容人数オーバーエラー
type ErrRoomCapacityOver struct {
	RoomID string
	Count  int
}

func (e ErrRoomCapacityOver) Error() string {
	return fmt.Sprintf("Room capacity over! - room: %s - count: %v", e.RoomID, e.Count)
}

// ErrRoomAlreadyEntry : 既入室ユーザーの入室リクエストエラー
type ErrRoomAlreadyEntry struct {
	RoomID string
}

func (e ErrRoomAlreadyEntry) Error() string {
	return fmt.Sprintf("Room already entry! - room: %s", e.RoomID)
}

// ErrNotRoomOwner : ルームの非所有者エラー
type ErrNotRoomOwner struct {
	RoomID string
}

func (e ErrNotRoomOwner) Error() string {
	return fmt.Sprintf("Not room owner! - room: %s", e.RoomID)
}

// ErrNotEntryRoom : 未入室ユーザーの退室リクエストエラー
type ErrNotEntryRoom struct {
	RoomID string
}

func (e ErrNotEntryRoom) Error() string {
	return fmt.Sprintf("Not entry room! - room: %s", e.RoomID)
}

// ErrDataBase : データベースエラー
type ErrDataBase struct {
	Detail interface{}
}

func (e ErrDataBase) Error() string {
	return fmt.Sprintf("Database error! - detail: %s", e.Detail)
}

// ErrNotFound : Not faoundエラー
type ErrNotFound struct {
	Detail interface{}
}

func (n ErrNotFound) Error() string {
	return fmt.Sprintf("Not found!")
}

// ErrRecordNotFound : レコードNotFaundエラー
type ErrRecordNotFound struct {
	Detail interface{}
}

func (e ErrRecordNotFound) Error() string {
	return fmt.Sprintf("Record not found! - detail: %s", e.Detail)
}

// ErrUnexpectedQueryProvider : 予想されていないプロバイダのクエリーエラー
type ErrUnexpectedQueryProvider struct {
	Provider string
}

func (p ErrUnexpectedQueryProvider) Error() string {
	return fmt.Sprintf("Unexpected query provider! - provider: %s", p.Provider)
}
