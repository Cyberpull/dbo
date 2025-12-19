package models

type PersonStatus struct {
	ID       uint64  `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	PersonID uint64  `json:"personId" gorm:"column:personId;index;not null"`
	Person   *Person `json:"person" gorm:"foreignKey:PersonID;OnUpdate:CASCADE;OnDelete:CASCADE"`
	Name     string  `json:"name" gorm:"column:name;index;size:100;not null"`
	Text     string  `json:"text" gorm:"column:text;index;size:200;not null"`
}
