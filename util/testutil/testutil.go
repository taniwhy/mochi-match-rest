package testutil

import (
	"time"

	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
)

// SetFakeUuID : モックデータ用のUUIDに置き換える
func SetFakeUuID(id string) {
	uuid.UuID = func() string {
		return id
	}
}

// SetFakeTime : モックデータ用の時刻に置き換える
func SetFakeTime(t time.Time) {
	clock.Now = func() time.Time {
		return t
	}
}
