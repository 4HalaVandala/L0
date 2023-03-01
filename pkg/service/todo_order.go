package service

import (
	"github.com/4halavandala/l0/internal/entity"
	"github.com/4halavandala/l0/pkg/cache"
	"github.com/4halavandala/l0/pkg/repository"
	"github.com/sirupsen/logrus"
)

type TodoOrderService struct {
	repo repository.TodoOrder
	tc   cache.TodoCache
}

func NewTodoOrderService(repo repository.TodoOrder, cache cache.TodoCache) *TodoOrderService {
	return &TodoOrderService{
		repo: repo,
		tc:   cache,
	}
}

func (s *TodoOrderService) Create(order entity.Model) (*entity.Model, error) {
	return s.repo.Create(order)
}

func (s *TodoOrderService) GetAll() (*[]entity.Model, error) {
	return s.repo.GetAll()
}

func (s *TodoOrderService) GetById(orderId string) (*entity.Model, error) {
	return s.repo.GetById(orderId)
}

func (s *TodoOrderService) InitCache() error {
	models, err := s.GetAll()
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	err = s.tc.AddAll(models)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	return nil
}

func (s *TodoOrderService) GetByIdFromCache(orderId string) (*entity.Model, error) {
	return s.tc.GetDataById(orderId)
}

func (s *TodoOrderService) GetAllFromCache() (*[]entity.Model, error) {
	return s.tc.GetAll()
}
