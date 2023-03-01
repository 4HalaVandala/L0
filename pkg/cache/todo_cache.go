package cache

import (
	"errors"
	"fmt"
	"github.com/4halavandala/l0/internal/entity"
)

type TodoCacheEntity struct {
	order *[]entity.Model
}

func NewTodoCacheEntity() *TodoCacheEntity {
	return &TodoCacheEntity{
		order: new([]entity.Model),
	}
}

func (tc *TodoCacheEntity) GetDataById(id string) (*entity.Model, error) {
	for _, order := range *tc.order {
		if order.Order_uid == id {
			return &order, nil
		}
	}
	err := fmt.Sprintf("Not found model with order_uid = %s", id)

	return nil, errors.New(err)
}

func (tc *TodoCacheEntity) AddAll(orders *[]entity.Model) error {
	tc.order = orders

	return nil
}

func (tc *TodoCacheEntity) GetAll() (*[]entity.Model, error) {
	return tc.order, nil
}
