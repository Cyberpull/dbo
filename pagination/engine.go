package pagination

import (
	"errors"
	"math"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Engine[T Setter[D], D any] struct{}

func (x *Engine[T, D]) Paginate(db *gorm.DB, options ...Options) (value T, err error) {
	opts := x.options(options...)

	db, err = x.model(db)

	if err != nil {
		return
	}

	db = db.Offset(-1).Limit(-1).Session(&gorm.Session{})

	data, total := []D{}, uint64(0)
	page, offset, from, limit := opts.parse()

	vtype := reflect.TypeOf(value)
	value = reflect.New(vtype).Interface().(T)
	value.SetCurrentPage(uint64(page))
	value.SetPerPage(limit)
	value.SetFrom(uint64(from))

	err = db.Offset(offset).Limit(limit).Find(&data).Error

	if err != nil {
		return
	}

	value.SetData(data)
	value.SetTo(uint64(offset + len(data)))

	tx := db.Offset(-1).Limit(-1).Clauses(clause.Select{
		Expression: clause.Expr{
			SQL:                "COUNT(?)",
			WithoutParentheses: true,
			Vars: []any{clause.Column{
				Table: clause.CurrentTable,
				Name:  clause.PrimaryKey,
			}},
		},
	})

	err = tx.Scan(&total).Error

	if err != nil {
		return
	}

	value.SetTotal(total)

	lastPage := math.Ceil(float64(total) / float64(limit))
	value.SetLastPage(uint64(lastPage))

	return
}

func (x *Engine[T, D]) options(options ...Options) Options {
	if len(options) > 0 {
		return options[0]
	}

	return Options{}
}

func (x *Engine[T, D]) model(db *gorm.DB) (tx *gorm.DB, err error) {
	var model D

	vType := reflect.TypeOf(model)

	var elem reflect.Type

	if elem = vType; elem.Kind() == reflect.Pointer {
		elem = elem.Elem()
	}

	if elem.Kind() != reflect.Struct {
		err = errors.New("model should be a struct")
		return
	}

	switch vType.Kind() {
	case reflect.Pointer:
		model = reflect.New(elem).Interface().(D)
		tx = db.Model(model)

	default:
		tx = db.Model(&model)
	}

	return
}

// =====================================

func init() {
	// var engine Engine[*Pagination[any]]
}
