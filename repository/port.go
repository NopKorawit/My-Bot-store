package repository

import (
	"store/model"
)

//Port
type GoodRepository interface {
	GetAllGoods() ([]model.Store, error)
	GetGoodsByType(types string) ([]model.Store, error)
	GetGoodsByCode(strcode string) (*model.Store, error)
	AddGoods(data model.StoreInput) (*model.Store, error)
	DeleteGood(strcode string) (*model.Store, error)
	UpdateGoodsByCode(strcode string, quantity int) (*model.Store, error)
}
