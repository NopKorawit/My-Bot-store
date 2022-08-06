package model

type Product struct {
	Code     int    `gorm:"size:5,index:idx_name,unique"`
	Type     string `gorm:"size:2,index:idx_name,unique"`
	Name     string `gorm:"size:30"`
	Quantity int    `gorm:"size:16"`
}

type ProductResponse struct {
	Code     string `gorm:"size:5"`
	Type     string `gorm:"size:2"`
	Name     string `gorm:"size:30"`
	Quantity int    `gorm:"size:16"`
}

type ProductInput struct {
	Type     string `gorm:"size:2"`
	Name     string `gorm:"size:30"`
	Quantity int    `gorm:"size:16"`
}

type MultiProduct struct {
	Code     string
	Quantity int
}
