package clauses

import (
	"gorm.io/gorm/clause"
)

type CTE interface {
	GetAlias() string
	GetExpression() clause.Expression
}

type With struct {
	CTE []CTE
}

func (x With) Name() string {
	return "WITH"
}

func (x With) MergeClause(cls *clause.Clause) {
	cls.Expression = x
}

func (x With) Build(builder clause.Builder) {
	for i, cte := range x.CTE {
		if i > 0 {
			builder.WriteString(", ")
		}

		alias, expr := cte.GetAlias(), cte.GetExpression()

		builder.WriteString(alias + " AS (")
		expr.Build(builder)
		builder.WriteString(")")
	}
}
