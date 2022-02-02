package controllers

import (
	"maranatha_web/internal/config"
	"maranatha_web/internal/repository"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DAO repository.DAO
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig, dao repository.DAO) *Repository {
	return &Repository{
		App: a,
		DAO: dao,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
