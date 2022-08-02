package service

import (
	"store/model"
)

//port
type GoodService interface {
	GetGoods() ([]model.StoreResponse, error)
	GetGoodsType(Type string) ([]model.StoreResponse, error)
	GetGood(code string) (*model.StoreResponse, error)
	AddGood(data model.StoreInput) (*model.StoreResponse, error)
	UpdateGood(code string, quantity int) (*model.StoreResponse, error)
	SellGood(code string, quantity int) (*model.StoreResponse, error)
	DelistGood(code string) (*model.StoreResponse, error)
}
