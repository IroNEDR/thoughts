package repository

import (
	"github.com/IroNEDR/thoughts/internals/config"
	"github.com/IroNEDR/thoughts/internals/db"
	"github.com/IroNEDR/thoughts/internals/models"
)

type Repository struct {
	app *config.AppConfig
	db  *db.Conn
}

func NewRepository(app *config.AppConfig, dbConn *db.Conn) *Repository {
	return &Repository{app, dbConn}
}

func (r *Repository) CreateUser(u models.User) error {
	return nil
}

func (r *Repository) GetUser(id uint) (models.User, error) {
	return models.User{}, nil
}

func (r *Repository) GetAllUsers() ([]models.User, error) {
	return nil, nil
}
