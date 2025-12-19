package ctes

import (
	"gorm.io/gorm/clause"
)

type PersonLocationCTE struct{}

func (x PersonLocationCTE) GetAlias() string {
	return "`pl`"
}

func (x PersonLocationCTE) GetExpression() clause.Expression {
	return clause.Expr{
		SQL:                "SELECT *, ROW_NUMBER() OVER (PARTITION BY `personId` ORDER BY `id` DESC) AS `row` FROM `person_location`",
		WithoutParentheses: true,
	}
}
