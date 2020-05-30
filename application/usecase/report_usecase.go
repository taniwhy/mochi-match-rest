package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models/dbmodel"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// ReportUsecase :
type ReportUsecase interface {
	FindAllReport() ([]*dbmodel.Report, error)
	InsertReport(report *dbmodel.Report) error
	DeleteReport(report *dbmodel.Report) error
}

type reportUsecase struct {
	reportRepository repository.ReportRepository
}

// NewReportUsecase :
func NewReportUsecase(rR repository.ReportRepository) ReportUsecase {
	return &reportUsecase{
		reportRepository: rR,
	}
}

func (rU reportUsecase) FindAllReport() ([]*dbmodel.Report, error) {
	chatposts, err := rU.reportRepository.FindAllReport()
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (rU reportUsecase) InsertReport(report *dbmodel.Report) error {
	err := rU.reportRepository.InsertReport(report)
	if err != nil {
		return err
	}
	return nil
}

func (rU reportUsecase) DeleteReport(report *dbmodel.Report) error {
	err := rU.reportRepository.DeleteReport(report)
	if err != nil {
		return err
	}
	return nil
}
