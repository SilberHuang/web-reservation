package dbrepo

import (
	"database/sql"

	"github.com/SilberHuang/web-reservation/internal/config"
	"github.com/SilberHuang/web-reservation/internal/repository"
)

type PostgresDatabaseRepo struct {
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &PostgresDatabaseRepo{
		App: a,
		DB: conn,
	}
}