package service

import (
	"github.com/No1ball/biocadTask/internal/entities"
	"github.com/No1ball/biocadTask/internal/repository"
)

type ApiFuncService struct {
	repo repository.ApiFunc
}

func NewApiFuncService(repo repository.ApiFunc) *ApiFuncService {
	return &ApiFuncService{repo: repo}
}

func (s *ApiFuncService) GetDataFromPage(offset, limit int) ([]entities.Entity, error) {
	return s.repo.GetDataFromPageRepo(offset, limit)
}
