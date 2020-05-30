package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
)

// ReportRepository :
type ReportRepository interface {
	FindAllReport() ([]*dbmodel.Report, error)
	InsertReport(report *dbmodel.Report) error
	DeleteReport(report *dbmodel.Report) error
}
