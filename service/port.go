package service

import (
	"store/model"
)

//port
type GoodService interface {
	GetGoods() ([]model.StoreResponse, error)
	GetGoodsType(Type string) ([]model.StoreResponse, error)
	AddGood(data model.StoreInput) (*model.StoreResponse, error)
	DelistGood(code string) (*model.StoreResponse, error)
}
