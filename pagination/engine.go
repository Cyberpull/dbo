package pagination

import (
	"errors"
	"math"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Engine[T Setter[D], D any] struct {
	clauseList []clause.Expression
}

func (x *Engine[T, D]) SetClauses(clauseList ...clause.Expression) {
	x.clauseList = clauseList
}

func (x *Engine[T, D]) Paginate(db *gorm.DB, options ...Options) (value T, err error) {
	opts := x.options(options...)

	db, err = x.model(db)

	if err != nil {
		return
	}

	db = db.Offset(-1).Limit(-1)

	data, total := []D{}, int64(0)
	page, offset, from, limit := opts.parse()

	if vt := reflect.TypeOf(value); vt.Kind() == reflect.Pointer {
		value = reflect.New(vt.Elem()).Interface().(T)
	}

	value.SetCurrentPage(uint64(page))
	value.SetPerPage(limit)
	value.SetFrom(uint64(from))

	err = db.Offset(offset).Limit(limit).Find(&data).Error

	if err != nil {
		return
	}

	value.SetData(data)
	value.SetTo(uint64(offset + len(data)))

	err = db.Offset(-1).Limit(-1).Count(&total).Error

	if err != nil {
		return
	}

	value.SetTotal(uint64(total))

	lastPage := math.Ceil(float64(total) / float64(limit))
	value.SetLastPage(uint64(lastPage))

	return
}

func (x *Engine[T, D]) applyClauses(db *gorm.DB) *gorm.DB {
	if len(x.clauseList) > 0 {
		db = db.Clauses(x.clauseList...)
	}

	return db
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
