package repository

import (
	"errors"
	"log"
	"store/model"

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

func (r goodRepositoryDB) CreateGoods(data model.StoreInput) (*model.Store, error) {
	good := model.Store{}
	result := r.db.Where("Name = ?", data.Name).Find(&good)
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
	return nil, errors.New("good already exists")
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
