package datatypes

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JSONArrayExpression struct {
	column   string
	path     string
	hasValue bool
	value    string
}

func JSONArrayQuery(column string, path string) *JSONArrayExpression {
	return &JSONArrayExpression{
		column: column,
		path:   path,
	}
}

func (query *JSONArrayExpression) HasValue(value string) *JSONArrayExpression {
	query.hasValue = true
	query.value = value
	return query
}

func (query *JSONArrayExpression) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		stmt.WriteString(fmt.Sprintf("jsonb_path_query_array(%v,", query.column))
		stmt.AddVar(stmt, query.path)
		stmt.WriteString(")")
		if query.hasValue {
			stmt.WriteString(" ? ")
			stmt.AddVar(stmt, query.value)
		}
	}
}
