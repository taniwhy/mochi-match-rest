package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type reportDatastore struct {
	db *gorm.DB
}

// NewReportDatastore : レポートデータストアの生成
func NewReportDatastore(db *gorm.DB) repository.IReportRepository {
	return &reportDatastore{db}
}

func (rD reportDatastore) FindAllReport() ([]*models.Report, error) {
	reports := []*models.Report{}

	err := rD.db.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (rD reportDatastore) InsertReport(report *models.Report) error {
	return rD.db.Create(report).Error
}

func (rD reportDatastore) DeleteReport(report *models.Report) error {
	err := rD.db.Take(&report).Error
	if err != nil {
		return err
	}
	return rD.db.Delete(report).Error
}
