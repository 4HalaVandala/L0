package repository

import (
	"encoding/json"
	"fmt"
	"github.com/4halavandala/l0/internal/entity"
	"github.com/jmoiron/sqlx"
)

type TodoOrderPostgres struct {
	db *sqlx.DB
}

func NewTodoOrderPostgres(db *sqlx.DB) *TodoOrderPostgres {
	return &TodoOrderPostgres{db: db}
}

func (r *TodoOrderPostgres) GetAll() (*[]entity.Model, error) {
	var lists []entity.Model
	query := fmt.Sprintf("SELECT * from %s", ordersTable)
	err := r.db.Select(&lists, query)

	return &lists, err
}

func (r *TodoOrderPostgres) GetById(orderId string) (*entity.Model, error) {
	var list entity.Model

	query := fmt.Sprintf(`SELECT * FROM %s o WHERE model->>'order_uid' = $1`, ordersTable)

	err := r.db.Get(&list, query, orderId)

	return &list, err
}

func (r *TodoOrderPostgres) Create(order entity.Model) (*entity.Model, error) {
	query := fmt.Sprintf("INSERT INTO %s (model) VALUES ($1) RETURNING model", ordersTable)
	jsonModel, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	var res entity.Model
	err = r.db.QueryRow(query, jsonModel).Scan(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
