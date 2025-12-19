package models

type PersonLocation struct {
	ID       uint64  `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	PersonID uint64  `json:"personId" gorm:"column:personId;index;not null"`
	Person   *Person `json:"person" gorm:"foreignKey:PersonID;OnUpdate:CASCADE;OnDelete:CASCADE"`
	Address  string  `json:"address" gorm:"column:address;index;size:200;not null"`
	City     string  `json:"city" gorm:"column:city;index;size:200;not null"`
	State    string  `json:"state" gorm:"column:state;index;size:200;not null"`
	Country  string  `json:"country" gorm:"column:country;index;size:200;not null"`
}
