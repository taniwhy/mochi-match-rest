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

func (d *reportDatastore) FindAll() ([]*models.Report, error) {
	reports := []*models.Report{}

	err := d.db.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (d *reportDatastore) Insert(report *models.Report) error {
	return d.db.Create(report).Error
}

func (d *reportDatastore) Delete(report *models.Report) error {
	err := d.db.Take(&report).Error
	if err != nil {
		return err
	}
	return d.db.Delete(report).Error
}
