package parser

import (
	"errors"
	"fmt"

	"github.com/seggga/he/internal/domain"
	"github.com/xwb1989/sqlparser"
)

type Parser struct {
	sql       string
	condition string
	parseStmt *sqlparser.Statement
	query     domain.Query
}

var (
	errStopParse = errors.New("stop parsing")
)

// NewParser creates new Parser with sql and ast fields
func NewParser(sql string) (*Parser, error) {
	parseStmt, err := sqlparser.Parse(sql)
	if err != nil {
		return nil, nil
	}
	return &Parser{
		sql:       sql,
		parseStmt: &parseStmt,
	}, nil
}

// ParseSelect parses sql and retreives columns from SEECT statement
func (p *Parser) Parse() error {
	selectStmt := make([]string, 0)
	fromStmt := make([]string, 0)
	var condition string

	visit := func(node sqlparser.SQLNode) (kontinue bool, err error) {
		switch node := node.(type) {
		case *sqlparser.Select:
			for _, v := range node.SelectExprs {
				selectStmt = append(selectStmt, sqlparser.String(v))
			}
			for _, v := range node.From {
				fromStmt = append(fromStmt, sqlparser.String(v))
			}

		case *sqlparser.Where:
			condition = sqlparser.String(node.Expr)
			return false, errStopParse
		}
		return true, nil
	}

	err := sqlparser.Walk(visit, *p.parseStmt)
	if err != nil && !errors.Is(err, errStopParse) {
		return fmt.Errorf("error parsing sql query, %w", err)
	}

	p.query.Select = selectStmt
	p.query.From = fromStmt
	p.condition = condition
	return nil
}

// GetSelect returns column names parsed from sql query (SELECT statement)
func (p Parser) GetSelect() []string {
	return p.query.Select
}

// GetFiles returns csv file names parsed from sql query (FROM statement)
func (p Parser) GetFiles() []string {
	return p.query.From
}

// GetCondition parses WHERE statement and produces slice of lexemmas
func (p Parser) GetCondition() {

}
