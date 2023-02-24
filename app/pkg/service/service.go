package service

import (
	"github.com/4halavandala/l0/app/internal/entity"
	"github.com/4halavandala/l0/app/pkg/repository"
)

type TodoOrder interface {
	Create() (int, error)
	GetAll() ([]entity.Order, error)
	GetById(orderId int) (entity.Order, error)
}

type CacheService interface {
	GetDataById(id int) (entity.Order, error)
}

type Service struct {
	TodoOrder
	CacheService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoOrder: NewTodoOrderService(repos.TodoOrder),
	}
}
