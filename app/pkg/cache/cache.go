package cache

import (
	"github.com/4halavandala/l0/app/internal/entity"
	"github.com/4halavandala/l0/app/pkg/repository"
)

type Cache struct {
	repo repository.CacheService
}

func (c *Cache) GetDataById(id int) (entity.Order, error) {
	return entity.Order{}, nil
}

func (c *Cache) initCache() []entity.Order {
	// select data from db
	// put data into map
	// save map
	//cache := make(map[int]string)

	return nil
}

func NewCache(repo repository.CacheService) *Cache {
	return &Cache{repo: repo}
}
