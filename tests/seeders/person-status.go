package seeders

import (
	"encoding/json"

	"github.com/Cyberpull/dbo/tests/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PersonStatus struct{}

func (x *PersonStatus) Name() string {
	return "Person Status"
}

func (x *PersonStatus) Handler(db *gorm.DB) (err error) {
	entries := []*models.PersonStatus{}

	data, err := FS.ReadFile("data/person-status.json")

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
			"name",
			"text",
		}),
	}).Create(entries).Error
}
