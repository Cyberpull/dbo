package ctes

import "gorm.io/gorm/clause"

type PersonStatusCTE struct{}

func (x PersonStatusCTE) GetAlias() string {
	return "`ps`"
}

func (x PersonStatusCTE) GetExpression() clause.Expression {
	return clause.Expr{
		SQL:                "SELECT *, ROW_NUMBER() OVER (PARTITION BY `personId` ORDER BY `id` DESC) AS `row` FROM `person_status`",
		WithoutParentheses: true,
	}
}
