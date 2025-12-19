package models

type House struct {
	ID          uint64  `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name        string  `json:"name" gorm:"column:name;index;size:100;not null"`
	Description string  `json:"description" gorm:"column:description;index;size:300;not null"`
	Price       float64 `json:"price" gorm:"column:price;index;type:float;not null"`
}
