package dbo

import (
	"time"

	"github.com/Cyberpull/dbo/clauses"
	"github.com/Cyberpull/dbo/scopes"

	"gorm.io/gorm"
)

type ConnectionCallback func(db *gorm.DB) *gorm.DB

type Connector interface {
	Connect(callback ...ConnectionCallback) (inst Instance, err error)
}

type xConnector struct {
	opts Options
}

func (x *xConnector) Connect(callback ...ConnectionCallback) (inst Instance, err error) {
	var db *gorm.DB
	var conn gorm.Dialector

	config := x.opts.getConfig()

	if conn, err = x.opts.getDialector(); err != nil {
		return
	}

	if db, err = gorm.Open(conn, config); err != nil {
		return
	}

	sqldb, err := db.DB()

	if err != nil {
		return
	}

	if x.opts.MaxOpenConns > 0 {
		sqldb.SetMaxOpenConns(x.opts.MaxOpenConns)
	}

	if x.opts.MaxIdleConns > 0 {
		sqldb.SetMaxIdleConns(x.opts.MaxIdleConns)
	}

	if x.opts.ConnMaxLifetime > 0 {
		duration := time.Millisecond * time.Duration(x.opts.ConnMaxLifetime)
		sqldb.SetConnMaxLifetime(duration)
	}

	if x.opts.ConnMaxIdleTime > 0 {
		duration := time.Millisecond * time.Duration(x.opts.ConnMaxIdleTime)
		sqldb.SetConnMaxIdleTime(duration)
	}

	clauses.RegisterPlugin(db)
	scopes.RegisterPlugin(db)

	switch x.opts.getDriver() {
	case DRIVER_PGSQL:
		err = db.Exec(`SET DEFAULT_TRANSACTION_ISOLATION TO SERIALIZABLE`).Error
		// SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;
		// SET DEFAULT_TRANSACTION_ISOLATION TO SERIALIZABLE;
	}

	db = x.call(db, callback...)

	inst = &xInstance{opts: x.opts, db: db}

	return
}

func (x *xConnector) call(db *gorm.DB, callbacks ...ConnectionCallback) *gorm.DB {
	for _, callback := range callbacks {
		db = callback(db)
	}

	return db
}

// ===============================

func NewConnector(opts Options) Connector {
	return &xConnector{opts: opts}
}
