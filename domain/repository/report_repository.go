//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

package repository

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
)

// IReportRepository : レポートのリポジトリ
type IReportRepository interface {
	FindAll() ([]*models.Report, error)
	Insert(report *models.Report) error
	Delete(report *models.Report) error
}
