package entity

import "fmt"

type keyOperator string

const (
	OperatorEqual   keyOperator = "="
	OperatorLess    keyOperator = "<"
	OperatorGreater keyOperator = ">"
)

type PrintRequest struct {
	Where []Where
}

type Where struct {
	Field    string
	Value    string
	Operator keyOperator
}

func (r *PrintRequest) BuildQuery(path string) []string {
	query := []string{path}
	for _, where := range r.Where {
		query = append(query, fmt.Sprintf(`?%s%s=%s`, where.Operator, where.Field, where.Value))
	}
	return query
}
