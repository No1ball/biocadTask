package repository

import (
	"github.com/No1ball/biocadTask/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ApiFunc interface {
	GetDataFromPageRepo(offset, limit int) ([]entities.Entity, error)
}

type TsvFile interface {
}
type Repository struct {
	ApiFunc
	TsvFile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		NewApiFuncPostgres(db),
		NewTsvFilePostgres(db),
	}
}
