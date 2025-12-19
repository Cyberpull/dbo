package models

import "gorm.io/gorm"

type Person struct {
	ID         uint64            `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name       string            `json:"name" gorm:"column:name;index;size:100;not null"`
	Cars       []*Car            `json:"cars" gorm:"many2many:person_car"`
	Houses     []*House          `json:"houses" gorm:"many2many:person_house"`
	StatusList []*PersonStatus   `json:"statusList" gorm:"foreignKey:PersonID"`
	Locations  []*PersonLocation `json:"locations" gorm:"foreignKey:PersonID"`
}

func (x *Person) PreloadCars(db *gorm.DB) *gorm.DB {
	return db
}

func (x *Person) PreloadHouses(db *gorm.DB) *gorm.DB {
	return db
}

func (x *Person) PreloadStatusList(db *gorm.DB) *gorm.DB {
	return db.Order("id DESC")
}

func (x *Person) PreloadLocations(db *gorm.DB) *gorm.DB {
	return db.Order("id DESC")
}
