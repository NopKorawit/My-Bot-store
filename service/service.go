package service

//Business logic in here

import (
	"Product/model"
	"Product/repository"
	"fmt"
	"log"
)

type ProductService struct {
	ProductRepo repository.ProductRepository //อ้างถึง interface
}

//constructor
func NewProductService(ProductRepo repository.ProductRepository) ProductService {
	return ProductService{ProductRepo: ProductRepo}
}

func (s ProductService) GetProducts() ([]model.ProductResponse, error) {
	Products, err := s.ProductRepo.GetAllProducts()
	if err != nil {
		log.Panic(model.ErrRepository)
		log.Println(err)
		return nil, err
	}
	qReponses := []model.ProductResponse{}
	for _, Product := range Products {
		qReponse := model.ProductResponse{
			Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
			Type:     Product.Type,
			Name:     Product.Name,
			Quantity: Product.Quantity,
		}
		qReponses = append(qReponses, qReponse)
	}

	return qReponses, nil
}

func (s ProductService) GetProductsType(Type string) ([]model.ProductResponse, error) {
	Products, err := s.ProductRepo.GetProductsByType(Type)
	if err != nil {
		log.Println(err)
		return nil, model.ErrRepository
	}

	qReponses := []model.ProductResponse{}
	for _, Product := range Products {
		qReponse := model.ProductResponse{
			Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
			Type:     Product.Type,
			Name:     Product.Name,
			Quantity: Product.Quantity,
		}
		qReponses = append(qReponses, qReponse)
	}

	return qReponses, nil
}

func (s ProductService) GetProduct(code string) (*model.ProductResponse, error) {
	Product, err := s.ProductRepo.GetProductsByCode(code)
	if err != nil {
		if err == model.ErrDuplicateROW {
			log.Println(err)
			return nil, err
		} else if err == model.ErrCodenotFound {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, model.ErrRepository
	}
	qReponse := model.ProductResponse{
		Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
		Type:     Product.Type,
		Name:     Product.Name,
		Quantity: Product.Quantity,
	}
	return &qReponse, nil
}

func (s ProductService) AddProduct(data model.ProductInput) (*model.ProductResponse, error) {
	Product, err := s.ProductRepo.AddProducts(data)
	if err != nil {
		if err == model.ErrDuplicateROW {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, model.ErrRepository
	}
	qReponse := model.ProductResponse{
		Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
		Type:     Product.Type,
		Name:     Product.Name,
		Quantity: Product.Quantity,
	}
	return &qReponse, nil
}

func (s ProductService) UpdateProduct(code string, quantity int) (*model.ProductResponse, error) {
	Product, err := s.ProductRepo.GetProductsByCode(code)
	if err != nil {
		if err == model.ErrDuplicateROW {
			log.Println(err)
			return nil, err
		} else if err == model.ErrCodenotFound {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, model.ErrRepository
	}
	Product.Quantity = quantity
	Product, err = s.ProductRepo.UpdateProductsByModel(Product)
	if err != nil {
		log.Println(err)
		return nil, model.ErrRepository
	}
	qReponse := model.ProductResponse{
		Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
		Type:     Product.Type,
		Name:     Product.Name,
		Quantity: Product.Quantity,
	}
	return &qReponse, nil
}

func (s ProductService) UpdateMultiProduct(code string, quantity int) (*model.ProductResponse, error) {
	Product, err := s.ProductRepo.GetProductsByCode(code)
	if err != nil {
		if err == model.ErrDuplicateROW {
			log.Println(err)
			return nil, err
		} else if err == model.ErrCodenotFound {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, model.ErrRepository
	}
	Product.Quantity = quantity
	Product, err = s.ProductRepo.UpdateProductsByModel(Product)
	if err != nil {
		log.Println(err)
		return nil, model.ErrRepository
	}
	qReponse := model.ProductResponse{
		Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
		Type:     Product.Type,
		Name:     Product.Name,
		Quantity: Product.Quantity,
	}
	return &qReponse, nil
}

func (s ProductService) SellProduct(code string, quantity int) (*model.ProductResponse, error) {
	Product, err := s.ProductRepo.GetProductsByCode(code)
	if err != nil {
		if err == model.ErrDuplicateROW {
			log.Println(err)
			return nil, err
		} else if err == model.ErrCodenotFound {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, model.ErrRepository
	}
	Product.Quantity = Product.Quantity - quantity
	if Product.Quantity < 0 {
		return nil, model.ErrProductNotEnough
	}
	new, err := s.ProductRepo.UpdateProductsByModel(Product)
	if err != nil {
		log.Println(err)
		return nil, model.ErrRepository
	}
	qReponse := model.ProductResponse{
		Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
		Type:     new.Type,
		Name:     new.Name,
		Quantity: new.Quantity,
	}
	return &qReponse, nil
}

func (s ProductService) SellMultiProduct(Products []model.MultiProduct) ([]model.ProductResponse, error) {
	var out []model.ProductResponse
	var in []model.Product
	for _, list := range Products {
		Product, err := s.ProductRepo.GetProductsByCode(list.Code)
		if err != nil {
			if err == model.ErrDuplicateROW {
				log.Println(err)
				return nil, err
			} else if err == model.ErrCodenotFound {
				log.Println(err)
				return nil, err
			}
			log.Println(err)
			return nil, model.ErrRepository
		}
		Product.Quantity = Product.Quantity - list.Quantity
		if Product.Quantity < 0 {
			qReponse := model.ProductResponse{
				Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
				Type:     Product.Type,
				Name:     Product.Name,
				Quantity: Product.Quantity,
			}
			out = append(out, qReponse)
		} else {
			in = append(in, *Product)
		}
	}
	if len(out) != 0 {
		return out, nil
	}
	for _, list2 := range Products {
		new, err := s.ProductRepo.UpdateProductsByCode2(list2.Code, list2.Quantity)
		if err != nil {
			log.Println(err)
			return nil, model.ErrRepository
		}
		qReponse := model.ProductResponse{
			Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
			Type:     new.Type,
			Name:     new.Name,
			Quantity: new.Quantity,
		}
	}

	// return &qReponse, nil
}

func (s ProductService) DelistProduct(code string) (*model.ProductResponse, error) {
	Product, err := s.ProductRepo.DeleteProduct(code)
	if err != nil {
		if err == model.ErrCodenotFound {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, model.ErrRepository
	}
	qReponse := model.ProductResponse{
		Code:     fmt.Sprintf("%v%d", Product.Type, Product.Code),
		Type:     Product.Type,
		Name:     Product.Name,
		Quantity: Product.Quantity,
	}
	return &qReponse, nil
}
