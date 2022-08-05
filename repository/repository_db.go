package repository

import (
	"Product/model"
	"log"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

//Adapter private
type ProductRepositoryDB struct {
	db *gorm.DB
}

//Constructor Public เพื่อ new instance
func NewProductRepositoryDB(db *gorm.DB) ProductRepository {
	return ProductRepositoryDB{db: db}
}

//buld all receiver function for interface
func (r ProductRepositoryDB) GetAllProducts() ([]model.Product, error) {
	Products := []model.Product{}
	err := r.db.Order("Code").Find(&Products).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return Products, nil
}

func (r ProductRepositoryDB) GetProductsByType(types string) ([]model.Product, error) {
	Products := []model.Product{}
	err := r.db.Where("Type = ?", types).Order("Code").Find(&Products).Error
	if err != nil {
		return nil, err
	}
	return Products, nil
}

func (r ProductRepositoryDB) GetProductsByCode(strcode string) (*model.Product, error) {
	Product := model.Product{}
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)
	result := r.db.Where("Code = ? AND Type = ?", code, Type).Find(&Product)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, model.ErrCodenotFound
	}
	if result.RowsAffected == 1 {
		return &Product, nil
	}
	return nil, model.ErrDuplicateROW
}

func (r ProductRepositoryDB) AddProducts(data model.ProductInput) (*model.Product, error) {
	Product := model.Product{}
	result := r.db.Where("Name = ? AND Type = ?", data.Name, data.Type).Find(&Product)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		newCode := r.generateCode(data.Type)
		Store := model.Product{
			Code:     newCode,
			Type:     data.Type,
			Name:     data.Name,
			Quantity: data.Quantity}
		r.db.Create(&Store)
		return &Store, nil
	}

	if result.RowsAffected == 1 {
		log.Println(Product)
		Product.Quantity = Product.Quantity + data.Quantity
		result = r.db.Where("Name = ? AND Type = ?", data.Name, data.Type).Save(&Product)
		if result.Error != nil {
			log.Println(result.Error)
			return nil, result.Error
		}
		return &Product, nil
	}
	return nil, model.ErrDuplicateROW
}

//not use
func (r ProductRepositoryDB) UpdateProductsByCode(strcode string, quantity int) (*model.Product, error) {
	Product, err := r.GetProductsByCode(strcode)
	if err != nil {
		return nil, err
	}
	Product.Quantity = quantity
	result := r.db.Where("Code = ? AND Type = ?", Product.Code, Product.Type).Save(&Product)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return Product, nil
}

func (r ProductRepositoryDB) UpdateProductsByCode2(strcode string, quantity int) (*model.Product, error) {
	Product := model.Product{}
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	types := strings.Trim(strcode, num)
	Product.Quantity = quantity
	result := r.db.Where("Code = ? AND Type = ?", code, types).Save(&Product)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return &Product, nil
}

func (r ProductRepositoryDB) UpdateProductsByModel(model *model.Product) (*model.Product, error) {
	Product := model
	result := r.db.Where("Code = ? AND Type = ?", model.Code, model.Type).Save(&Product)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	return Product, nil
}

func (r ProductRepositoryDB) generateCode(types string) (NewCode int) {
	Product := model.Product{}
	result := r.db.Where("Type=?", types).Limit(1).Order("Code desc").Find(&Product)
	log.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		NewCode = 1
		return NewCode
	}
	NewCode = Product.Code + 1
	return NewCode
}

func (r ProductRepositoryDB) DeleteProduct(strcode string) (*model.Product, error) {
	var Product model.Product
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	types := strings.Trim(strcode, num)
	result := r.db.Where("Code = ? AND Type = ?", code, types).Find(&Product)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, model.ErrCodenotFound
	}
	r.db.Where("Code = ? AND Type = ?", code, types).Delete(&Product)
	return &Product, nil
}
