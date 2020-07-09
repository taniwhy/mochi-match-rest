package models

import (
	"time"

	"github.com/taniwhy/mochi-match-rest/util/clock"
	"github.com/taniwhy/mochi-match-rest/util/uuid"
)

// Report : reportテーブルモデル
type Report struct {
	ReportID         string
	CallerID         string
	VaiolatorID      string
	ReportedRoomID   string
	VaiolationDetail string
	CreatedAt        time.Time
}

// NewReport :
func NewReport(callerID, violatorID, roomID, detail string) (*Report, error) {
	return &Report{
		ReportID:         uuid.UuID(),
		CallerID:         callerID,
		VaiolatorID:      violatorID,
		ReportedRoomID:   roomID,
		VaiolationDetail: detail,
		CreatedAt:        clock.Now(),
	}, nil
}
