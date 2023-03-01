package service

import (
	"github.com/4halavandala/l0/internal/entity"
	"github.com/4halavandala/l0/pkg/cache"
	"github.com/4halavandala/l0/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type TodoOrder interface {
	Create(order entity.Model) (*entity.Model, error)
	GetAll() (*[]entity.Model, error)
	GetById(orderId string) (*entity.Model, error)

	InitCache() error
	GetByIdFromCache(orderId string) (*entity.Model, error)
	GetAllFromCache() (*[]entity.Model, error)
}

type Service struct {
	TodoOrder
}

func NewService(repos *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		TodoOrder: NewTodoOrderService(repos.TodoOrder, cache.TodoCache),
	}
}
