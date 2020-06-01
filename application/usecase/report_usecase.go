package usecase

import (
	"github.com/taniwhy/mochi-match-rest/domain/models"
	"github.com/taniwhy/mochi-match-rest/domain/repository"
)

// ReportUsecase :
type ReportUsecase interface {
	FindAllReport() ([]*models.Report, error)
	InsertReport(report *models.Report) error
	DeleteReport(report *models.Report) error
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

func (rU reportUsecase) FindAllReport() ([]*models.Report, error) {
	chatposts, err := rU.reportRepository.FindAllReport()
	if err != nil {
		return nil, err
	}
	return chatposts, nil
}

func (rU reportUsecase) InsertReport(report *models.Report) error {
	err := rU.reportRepository.InsertReport(report)
	if err != nil {
		return err
	}
	return nil
}

func (rU reportUsecase) DeleteReport(report *models.Report) error {
	err := rU.reportRepository.DeleteReport(report)
	if err != nil {
		return err
	}
	return nil
}
