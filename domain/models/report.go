package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/taniwhy/mochi-match-rest/domain/errors"
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
func NewReport(cID, vID, rID, vD string) (*Report, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.ErrGenerateID{}
	}
	return &Report{
		ReportID:         id.String(),
		CallerID:         cID,
		VaiolatorID:      vID,
		ReportedRoomID:   rID,
		VaiolationDetail: vD,
		CreatedAt:        time.Now(),
	}, nil
}
