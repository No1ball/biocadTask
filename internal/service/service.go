package service

import (
	"github.com/No1ball/biocadTask/internal/entities"
	"github.com/No1ball/biocadTask/internal/repository"
)

type ApiFunc interface {
	GetDataFromPage(offset, limit int) ([]entities.Entity, error)
}
type TsvFileFunc interface {
}
type Service struct {
	ApiFunc
	TsvFileFunc
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		NewApiFuncService(repo.ApiFunc),
		NewTsvFileService(repo.TsvFile),
	}
}
