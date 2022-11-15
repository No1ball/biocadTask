package repository

import "github.com/jmoiron/sqlx"

type TsvFilePostgres struct {
	db *sqlx.DB
}

func NewTsvFilePostgres(db *sqlx.DB) *TsvFilePostgres {
	return &TsvFilePostgres{db: db}
}
