package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// ReportRepository :
type ReportRepository interface {
	FindAllReport() ([]*models.Report, error)
	InsertReport(report *models.Report) error
	DeleteReport(report *models.Report) error
}
