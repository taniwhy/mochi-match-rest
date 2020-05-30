package dbmodel

import (
	"time"
)

// Report : reportテーブルモデル
type Report struct {
	ID               int
	Caller           int
	Vaiolator        int
	VaiolationDetail int
	ReportedRoom     int
	CreatedAt        time.Time
}
