package repository

import (
	"github.com/4halavandala/l0/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TodoOrder interface {
	Create(order entity.Model) (*entity.Model, error)
	GetAll() (*[]entity.Model, error)
	GetById(orderId string) (*entity.Model, error)
}

type Repository struct {
	TodoOrder
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoOrder: NewTodoOrderPostgres(db),
	}
}
