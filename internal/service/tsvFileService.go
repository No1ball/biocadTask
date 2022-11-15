package service

import "github.com/No1ball/biocadTask/internal/repository"

type TsvFileService struct {
	repo repository.TsvFile
}

func NewTsvFileService(repo repository.TsvFile) *TsvFileService {
	return &TsvFileService{repo: repo}
}
