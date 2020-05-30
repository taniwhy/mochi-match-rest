package datastore

import (
	"github.com/jinzhu/gorm"
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

type reportDatastore struct {
	db *gorm.DB
}

// NewReportDatastore : .
func NewReportDatastore(db *gorm.DB) repository.ReportRepository {
	return &reportDatastore{db}
}

func (rD reportDatastore) FindAllReport() ([]*dbmodel.Report, error) {
	reports := []*dbmodel.Report{}

	err := rD.db.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (rD reportDatastore) InsertReport(report *dbmodel.Report) error {
	return rD.db.Create(report).Error
}

func (rD reportDatastore) DeleteReport(report *dbmodel.Report) error {
	err := rD.db.Take(&report).Error
	if err != nil {
		return err
	}
	return rD.db.Delete(report).Error
}
