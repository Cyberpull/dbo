package dbo

import "gorm.io/gorm"

func NewSession(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{
		NewDB: true,
	})
}
