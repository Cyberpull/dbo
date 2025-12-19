package seeders

import (
	"encoding/json"

	"github.com/Cyberpull/dbo/tests/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Person struct{}

func (x *Person) Name() string {
	return "Persons"
}

func (x *Person) Handler(db *gorm.DB) (err error) {
	entries := []*models.Person{}

	data, err := FS.ReadFile("data/persons.json")

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &entries)

	if err != nil {
		return
	}

	return db.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Omit(clause.Associations).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name"}),
		}).Create(entries).Error

		if err != nil {
			return
		}

		for _, person := range entries {
			err = tx.Model(person).Association("Cars").Replace(person.Cars)

			if err != nil {
				return
			}

			err = tx.Model(person).Association("Houses").Replace(person.Houses)

			if err != nil {
				return
			}
		}

		return
	})
}
