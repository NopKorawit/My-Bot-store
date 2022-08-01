package service

//Business logic in here

import (
	"fmt"
	"log"
	"store/model"
	"store/repository"
)

type goodService struct {
	goodRepo repository.GoodRepository //อ้างถึง interface
}

//constructor
func NewGoodService(goodRepo repository.GoodRepository) GoodService {
	return goodService{goodRepo: goodRepo}
}

func (s goodService) GetGoods() ([]model.StoreResponse, error) {
	goods, err := s.goodRepo.GetAllGoods()
	if err != nil {
		log.Panic(model.ErrRepository)
		log.Println(err)
		return nil, err
	}
	qReponses := []model.StoreResponse{}
	for _, good := range goods {
		qReponse := model.StoreResponse{
			Code:     fmt.Sprintf("%v%03d", good.Type, good.Code),
			Type:     good.Type,
			Name:     good.Name,
			Quantity: good.Quantity,
		}
		qReponses = append(qReponses, qReponse)
	}

	return qReponses, nil
}

func (s goodService) GetGoodsType(Type string) ([]model.StoreResponse, error) {
	goods, err := s.goodRepo.GetGoodsByType(Type)
	if err != nil {
		log.Println(err)
		return nil, model.ErrRepository
	}

	qReponses := []model.StoreResponse{}
	for _, good := range goods {
		qReponse := model.StoreResponse{
			Code:     fmt.Sprintf("%v%03d", good.Type, good.Code),
			Type:     good.Type,
			Name:     good.Name,
			Quantity: good.Quantity,
		}
		qReponses = append(qReponses, qReponse)
	}

	return qReponses, nil
}

func (s goodService) AddGood(data model.StoreInput) (*model.StoreResponse, error) {
	good, err := s.goodRepo.CreateGoods(data)
	if err != nil {
		if err == model.ErrGoodAlreadyExists {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, model.ErrRepository
	} else {
		qReponse := model.StoreResponse{
			Code:     fmt.Sprintf("%v%03d", good.Type, good.Code),
			Type:     good.Type,
			Name:     good.Name,
			Quantity: good.Quantity,
		}
		return &qReponse, nil
	}
}

func (s goodService) DelistGood(code string) (*model.StoreResponse, error) {
	good, err := s.goodRepo.DeleteGood(code)
	if err != nil {
		if err == model.ErrCodenotFound {
			log.Println(err)
			return nil, err
		}
		log.Println(err)
		return nil, model.ErrRepository
	}
	qReponse := model.StoreResponse{
		Code:     fmt.Sprintf("%v%03d", good.Type, good.Code),
		Type:     good.Type,
		Name:     good.Name,
		Quantity: good.Quantity,
	}
	return &qReponse, nil
}
