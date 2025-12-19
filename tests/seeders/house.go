package seeders

import (
	"github.com/Cyberpull/dbo/tests/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type House struct{}

func (x *House) Name() string {
	return "Houses"
}

func (x *House) Handler(db *gorm.DB) (err error) {
	entries := []*models.House{
		{ID: 1, Name: "House 1", Description: "8 Bedrooms, large sitting room, swimming pool", Price: 3500000},
		{ID: 2, Name: "House 2", Description: "4 Bedrooms, large sitting room, heater", Price: 3800000},
		{ID: 3, Name: "House 3", Description: "6 Bedrooms, large sitting room, swimming pool", Price: 3600000},
		{ID: 4, Name: "House 4", Description: "16 Bedrooms, large sitting room, swimming pool, heater", Price: 2600000},
		{ID: 5, Name: "House 5", Description: "5 Bedrooms, large sitting room, heater", Price: 3300000},
		{ID: 6, Name: "House 6", Description: "12 Bedrooms, large sitting room, swimming pool, heater", Price: 4100000},
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
