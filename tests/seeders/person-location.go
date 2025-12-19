package seeders

import (
	"encoding/json"

	"github.com/Cyberpull/dbo/tests/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PersonLocation struct{}

func (x *PersonLocation) Name() string {
	return "Person Locations"
}

func (x *PersonLocation) Handler(db *gorm.DB) (err error) {
	entries := []*models.PersonLocation{}

	data, err := FS.ReadFile("data/person-locations.json")

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &entries)

	if err != nil {
		return
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"personId",
			"address",
			"city",
			"state",
			"country",
		}),
	}).Create(entries).Error
}
