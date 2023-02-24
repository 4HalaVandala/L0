package service

import "github.com/4halavandala/l0/app/internal/entity"

func GetDataById(id int) entity.Order {
	return entity.Order{}
}

func initCache() []entity.Order {
	// select data from db
	// put data into map
	// save map
	cache := make(map[int]string)

	return nil
}
