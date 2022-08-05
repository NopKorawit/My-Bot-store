package repository

import (
	"Product/model"
)

//Port
type ProductRepository interface {
	GetAllProducts() ([]model.Product, error)
	GetProductsByType(types string) ([]model.Product, error)
	GetProductsByCode(strcode string) (*model.Product, error)
	AddProducts(data model.ProductInput) (*model.Product, error)
	DeleteProduct(strcode string) (*model.Product, error)
	UpdateProductsByCode(strcode string, quantity int) (*model.Product, error)
	UpdateProductsByCode2(strcode string, quantity int) (*model.Product, error)
	UpdateProductsByModel(model *model.Product) (*model.Product, error)
}
