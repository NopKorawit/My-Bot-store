package repository

import (
	"log"
	"store/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

//Adapter private
type goodRepositoryDB struct {
	db *gorm.DB
}

//Constructor Public เพื่อ new instance
func NewGoodRepositoryDB(db *gorm.DB) GoodRepository {
	return goodRepositoryDB{db: db}
}

//buld all receiver function for interface

func (r goodRepositoryDB) GetAllGoods() ([]model.Store, error) {
	goods := []model.Store{}
	err := r.db.Order("Code").Find(&goods).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return goods, nil
}

func (r goodRepositoryDB) GetGoodsByType(types string) ([]model.Store, error) {
	goods := []model.Store{}
	err := r.db.Where("Type = ?", types).Order("Code").Find(&goods).Error
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (r goodRepositoryDB) AddGoods(data model.StoreInput) (*model.Store, error) {
	good := model.Store{}
	result := r.db.Where("Name = ? AND Type = ?", data.Name, data.Type).Find(&good)
	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		newCode := r.generateCode(data.Type)
		Store := model.Store{
			Code:     newCode,
			Type:     data.Type,
			Name:     data.Name,
			Quantity: data.Quantity}
		r.db.Create(&Store)
		return &Store, nil
	}

	if result.RowsAffected == 1 {
		log.Println("เข้า1")
		log.Println(good)
		good.Quantity = data.Quantity
		result = r.db.Where("Name = ? AND Type = ?", data.Name, data.Type).Save(&good)
		if result.Error != nil {
			log.Println(result.Error)
			return nil, result.Error
		}
		return &good, nil
	}
	return nil, model.ErrDuplicateROW
}

func (r goodRepositoryDB) UpdateGoodsByCode(strcode string, quantity int) (*model.Store, error) {
	good := model.Store{}
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)
	result := r.db.Where("Code = ? AND Type = ?", code, Type).Find(&good)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, model.ErrCodenotFound
	}
	if result.RowsAffected == 1 {
		good.Quantity = quantity
		result = r.db.Where("Code = ? AND Type = ?", code, Type).Save(&good)
		if result.Error != nil {
			log.Println(result.Error)
			return nil, result.Error
		}
		return &good, nil
	}
	return nil, model.ErrDuplicateROW
}

func (r goodRepositoryDB) GetGoodsByCode(strcode string) (*model.Store, error) {
	good := model.Store{}
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	Type := strings.Trim(strcode, num)
	result := r.db.Where("Code = ? AND Type = ?", code, Type).Find(&good)
	if result.Error != nil {
		return nil, result.Error
	}
	print(result.RowsAffected)
	if result.RowsAffected == 0 {
		return nil, model.ErrCodenotFound
	}
	if result.RowsAffected == 1 {
		return &good, nil
	}
	return nil, model.ErrDuplicateROW
}

func (r goodRepositoryDB) generateCode(types string) (NewCode int) {
	good := model.Store{}
	result := r.db.Where("Type=?", types).Limit(1).Order("Code desc").Find(&good)
	log.Println(result.RowsAffected)
	if result.RowsAffected == 0 {
		NewCode = 1
		return NewCode
	}
	NewCode = good.Code + 1
	return NewCode
}

func (r goodRepositoryDB) DeleteGood(strcode string) (*model.Store, error) {
	var good model.Store
	num := strings.TrimLeft(strcode, "ABCD")
	code, _ := strconv.Atoi(num)
	types := strings.Trim(strcode, num)
	result := r.db.Where("Code = ? AND Type = ?", code, types).Find(&good)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, model.ErrCodenotFound
	}
	r.db.Where("Code = ? AND Type = ?", code, types).Delete(&good)
	return &good, nil
}
