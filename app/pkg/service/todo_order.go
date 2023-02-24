package service

import (
	"github.com/4halavandala/l0/app/internal/entity"
	"github.com/4halavandala/l0/app/pkg/repository"
)

type TodoOrderService struct {
	repo repository.TodoOrder
}

func (s *TodoOrderService) Create() (int, error) {
	return s.repo.Create()
}

func (s *TodoOrderService) GetAll() ([]entity.Order, error) {
	return s.repo.GetAll()
}

func (s *TodoOrderService) GetById(orderId int) (entity.Order, error) {
	return s.repo.GetById(orderId)
}

func NewTodoOrderService(repo repository.TodoOrder) *TodoOrderService {
	return &TodoOrderService{repo: repo}
}
