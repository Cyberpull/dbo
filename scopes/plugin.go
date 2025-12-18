package scopes

import (
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type Scope = func(db *gorm.DB) *gorm.DB

type xPluginCallback func(db *gorm.DB)

type xPlugin struct{}

func (x *xPlugin) Name() string {
	return "dbo"
}

func (x *xPlugin) Initialize(db *gorm.DB) (err error) {
	db.Callback().Query().Before("*").Register("dbo:before_query", x.onBeforeQuery())
	db.Callback().Query().After("*").Register("dbo:after_query", x.onAfterQuery())

	return
}

// ==================================

func (x *xPlugin) onBeforeQuery() xPluginCallback {
	return func(db *gorm.DB) {
		defer func() {
			_ = recover()
		}()

		model := reflect.New(db.Statement.Schema.ModelType)

		// Process Scopes
		for i := 0; i < model.NumMethod(); i++ {
			name := model.Type().Method(i).Name

			if strings.HasPrefix(name, "Scope") {
				method, ok := model.Method(i).Interface().(Scope)

				if ok {
					db = method(db)
				}
			}
		}

		// Process Tags
		for _, field := range db.Statement.Schema.Fields {
			method := model.MethodByName("Preload" + field.Name)

			if method.IsValid() && !method.IsZero() {
				db = db.Preload(field.Name, method.Interface())
			}
		}
	}
}

func (x *xPlugin) onAfterQuery() xPluginCallback {
	return func(db *gorm.DB) {
		// for _, field := range db.Statement.Schema.Fields {
		// 	tag := x.getFieldTag(field)
		// }
	}
}

// ===============================

func RegisterPlugin(db *gorm.DB) (err error) {
	return db.Use(&xPlugin{})
}
