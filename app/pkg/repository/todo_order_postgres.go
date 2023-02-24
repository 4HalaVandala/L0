package repository

import (
	"fmt"
	"github.com/4halavandala/l0/app/internal/entity"
	"github.com/jmoiron/sqlx"
)

const (
	ordersTable = "orders"
)

type TodoOrderPostgres struct {
	db *sqlx.DB
}

func (r *TodoOrderPostgres) GetAll() ([]entity.Order, error) {
	var lists []entity.Order
	query := fmt.Sprintf(`SELECT * from %s`, ordersTable)
	err := r.db.Select(&lists, query)

	return lists, err
}

func (r *TodoOrderPostgres) GetById(orderId int) (entity.Order, error) {
	var list entity.Order

	query := fmt.Sprintf(`SELECT * FROM %s o WHERE o.id = $1`, ordersTable)

	err := r.db.Get(&list, query, orderId)

	return list, err
}

func (r *TodoOrderPostgres) Create() (int, error) {
	return 1, nil
}

func NewTodoOrderPostgres(db *sqlx.DB) *TodoOrderPostgres {
	return &TodoOrderPostgres{db: db}
}
