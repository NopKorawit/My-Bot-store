package service

import (
	"Product/model"
)

//port
type ProductService interface {
	GetProducts() ([]model.ProductResponse, error)
	GetProductsType(Type string) ([]model.ProductResponse, error)
	GetProduct(code string) (*model.ProductResponse, error)
	AddProduct(data model.ProductInput) (*model.ProductResponse, error)
	UpdateProduct(code string, quantity int) (*model.ProductResponse, error)
	SellProduct(code string, quantity int) (*model.ProductResponse, error)
	DelistProduct(code string) (*model.ProductResponse, error)
}
