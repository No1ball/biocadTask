package repository

import (
	"fmt"
	"github.com/No1ball/biocadTask/internal/entities"
	"github.com/jmoiron/sqlx"
)

type ApiFuncPostgres struct {
	db *sqlx.DB
}

func NewApiFuncPostgres(db *sqlx.DB) *ApiFuncPostgres {
	return &ApiFuncPostgres{db: db}
}

func (r *ApiFuncPostgres) GetDataFromPageRepo(offset, limit int) ([]entities.Entity, error) {
	var data []entities.Entity

	query := fmt.Sprintf("SELECT * FROM %s OFFSET $1 LIMIT $2", dataTable)
	err := r.db.Select(&data, query, offset, limit)

	return data, err
}
