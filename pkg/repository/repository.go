package repository

import "github.com/jmoiron/sqlx"

type Data interface {
	CollectData(username string, chatid int64, message string, answer []string) error
	GetNumberOfUsers() (int64, error)
}

type Repositroy struct {
	Data
}

func NewRepository(db *sqlx.DB) *Repositroy {
	return &Repositroy{
		Data: NewDataPostgres(db),
	}
}
