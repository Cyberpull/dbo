package seeders

import (
	"github.com/Cyberpull/dbo/tests/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Car struct{}

func (x *Car) Name() string {
	return "Cars"
}

func (x *Car) Handler(db *gorm.DB) (err error) {
	entries := []*models.Car{
		{ID: 1, Name: "Toyota", Description: "Toyota Corolla", Price: 3500000},
		{ID: 2, Name: "Toyota", Description: "Toyota Camry", Price: 3800000},
		{ID: 3, Name: "Honda", Description: "Honda Accord", Price: 3600000},
		{ID: 4, Name: "Peugeot", Description: "Peugeot 504", Price: 2600000},
		{ID: 5, Name: "Peugeot", Description: "Peugeot 360", Price: 3300000},
		{ID: 6, Name: "Hyudai", Description: "Hyudai 2024", Price: 4100000},
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"name",
			"description",
			"price",
		}),
	}).Create(entries).Error
}
