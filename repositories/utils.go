package repositories

import (
	"fmt"

	"template-service/models"
)

// AddLimitOffsetClause conditionally adds a LIMIT and OFFSET clause to a query
func AddLimitOffsetClause(query string, parameterValues []interface{}, pagingParams models.PagingParameters) (string, []interface{}) {

	// Conditionally add an OFFSET clause
	if pagingParams.Offset > 0 {
		parameterValues = append(parameterValues, pagingParams.Offset)
		query += fmt.Sprintf(" OFFSET $%d", len(parameterValues))
	}

	// Conditionally add an LIMIT clause
	if pagingParams.Limit > 0 {
		parameterValues = append(parameterValues, pagingParams.Limit)
		query += fmt.Sprintf(" LIMIT $%d", len(parameterValues))
	}

	return query, parameterValues
}
