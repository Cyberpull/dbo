package clauses

import "gorm.io/gorm"

type xPlugin struct{}

func (x *xPlugin) Name() string {
	return "dbo-clauses"
}

func (x *xPlugin) Initialize(db *gorm.DB) (err error) {
	clauses := []string{"WITH"}
	clauses = append(clauses, db.Callback().Query().Clauses...)

	db.Callback().Query().Clauses = clauses

	return
}

// ==========================================

func RegisterPlugin(db *gorm.DB) (err error) {
	return db.Use(&xPlugin{})
}
