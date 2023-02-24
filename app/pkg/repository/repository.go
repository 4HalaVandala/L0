package repository

import (
	"github.com/4halavandala/l0/app/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TodoOrder interface {
	Create() (int, error)
	GetAll() ([]entity.Order, error)
	GetById(orderId int) (entity.Order, error)
}

type CacheService interface {
	GetDataById(id int) (entity.Order, error)
}

type Repository struct {
	TodoOrder
	CacheService
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoOrder: NewTodoOrderPostgres(db),
	}
}
