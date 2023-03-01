package cache

import (
	"github.com/4halavandala/l0/internal/entity"
)

type Cache struct {
	TodoCache
}

type TodoCache interface {
	GetDataById(id string) (*entity.Model, error)
	GetAll() (*[]entity.Model, error)
	AddAll(orders *[]entity.Model) error
}

func NewCache() *Cache {
	return &Cache{
		TodoCache: NewTodoCacheEntity(),
	}
}
