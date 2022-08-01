package model

type Store struct {
	Code     int    `gorm:"size:5"`
	Type     string `gorm:"size:2"`
	Name     string `gorm:"size:30"`
	Quantity string `gorm:"size:16"`
}

type StoreResponse struct {
	Code     string `gorm:"size:5"`
	Type     string `gorm:"size:2"`
	Name     string `gorm:"size:30"`
	Quantity string `gorm:"size:16"`
}

type StoreInput struct {
	Type     string `gorm:"size:2"`
	Name     string `gorm:"size:30"`
	Quantity string `gorm:"size:16"`
}
