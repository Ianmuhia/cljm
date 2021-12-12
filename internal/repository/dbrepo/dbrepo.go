package dbrepo

import (
	"database/sql"

	"bitbucket.org/wycemiro/cljm.git/internal/config"
	"bitbucket.org/wycemiro/cljm.git/internal/repository"
)

type PostgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &PostgresDBRepo{
		App: a,
		DB:  conn,
	}
}
