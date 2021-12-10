package parser

import (
	"github.com/seggga/he/internal/domain"
	"github.com/xwb1989/sqlparser"
)

type Parser struct {
	sql   string
	ast   *sqlparser.Statement
	query domain.Query
}

func NewParser(sql string) (*Parser, error) {
	ast, err := sqlparser.Parse(sql)
	if err != nil {
		return nil, nil
	}
	return &Parser{
		sql: sql,
		ast: &ast,
	}, nil
}

func (p *Parser) ParseSelect() {
	// sql := `select one, two, three from file_one.csv, file_two.csv where one>10 and two="hi there"`
	// stmt, _ := sqlparser.Parse(p.sql)

	selectStmt := make([]string, 0)
	visitSelect := func(node sqlparser.SQLNode) (kontinue bool, err error) {
		if sel, ok := node.(*sqlparser.Select); ok {
			for _, v := range sel.SelectExprs {
				selectStmt = append(selectStmt, sqlparser.String(v))
			}
			return false, nil
		}
		return true, nil
	}
	sqlparser.Walk(visitSelect, *p.ast)

	p.query.Select = selectStmt
}
