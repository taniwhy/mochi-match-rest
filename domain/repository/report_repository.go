package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// ReportRepository : レポートのリポジトリ
type ReportRepository interface {
	FindAllReport() ([]*models.Report, error)
	InsertReport(report *models.Report) error
	DeleteReport(report *models.Report) error
}
