package repository

import (
	"store/model"
)

//Port
type GoodRepository interface {
	GetAllGoods() ([]model.Store, error)
	GetGoodsByType(types string) ([]model.Store, error)
	CreateGoods(data model.StoreInput) (*model.Store, error)
}
