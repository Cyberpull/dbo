package dbo

import (
	"context"
	"time"

	"github.com/Cyberpull/dbo/scopes"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DB struct {
	DB      *gorm.DB
	Scopes  []scopes.Scope
	Clauses []clause.Expression
	Timeout time.Duration
}

func (x *DB) newContext() (ctx context.Context, cancel context.CancelFunc) {
	if x.Timeout > 0 {
		return context.WithTimeout(context.Background(), x.Timeout)
	}

	return context.WithCancel(context.Background())
}
