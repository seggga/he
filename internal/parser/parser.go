package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/seggga/he/internal/domain"
	"github.com/xwb1989/sqlparser"
)

type Parser struct {
	condition   string // TODO прояснить обоснованность
	ast         *sqlparser.Statement
	parsedQuery domain.ParsedQuery
}

var (
	errStopParse   = errors.New("stop parsing")
	errWrongSelect = errors.New("wrong query, SELECT is empty")
	errWrongFrom   = errors.New("wrong query, FROM is empty")
)

// NewParser creates new Parser
func NewParser() Parser {
	return Parser{}
}

// ParseSelect parses sql and retreives columns from SEECT statement
func (p *Parser) Parse(sql string) error {

	ast, err := sqlparser.Parse(sql)
	if err != nil {
		return fmt.Errorf("error parsing sql-query, %w", err)
	}

	selectStmt := make([]string, 0)
	fromStmt := make([]string, 0)
	condition := ""

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

	err = sqlparser.Walk(visit, ast)
	if err != nil && !errors.Is(err, errStopParse) {
		return fmt.Errorf("error parsing sql query, %w", err)
	}

	if len(selectStmt) == 0 {
		return errWrongSelect
	}
	if len(fromStmt) == 0 {
		return errWrongFrom
	}

	p.parsedQuery.Select = selectStmt
	p.parsedQuery.Files = fromStmt
	p.parsedQuery.Where = parseCondition(condition)

	return nil
}

// GetSelect returns column names parsed from sql query (SELECT statement)
func (p Parser) GetSelect() []string {
	return p.parsedQuery.Select
}

// GetFiles returns csv file names parsed from sql query (FROM statement)
func (p Parser) GetFiles() []string {
	return p.parsedQuery.Files
}

// GetCondition parses WHERE statement and produces slice of lexemmas
func (p Parser) GetCondition() []domain.Token {
	return nil
}

// parse condition by sequential token scanning
func parseCondition(condition string) []domain.Token {
	where := make([]domain.Token, 0)

	r := strings.NewReader(condition)
	tokenizer := sqlparser.NewTokenizer(r)
	for {
		i, b := tokenizer.Scan()
		if i == 0 {
			break
		}
		if isValidToken(i) {
			token := composeToken(i, b)
			where = append(where, token)
		}
	}

	return where
}

// token: {token, lexema, tokenType, priority}
var validTokens = map[int]domain.Token{
	40:    {40, []byte("("), "operator", 0},
	41:    {41, []byte(")"), "operator", 0},
	60:    {60, []byte("<"), "operator", 0},
	61:    {61, []byte("="), "operator", 0},
	62:    {62, []byte(">"), "operator", 0},
	57418: {57418, []byte("<="), "operator", 0},
	57419: {57419, []byte(">="), "operator", 0},
	57420: {57420, []byte("!="), "operator", 0},
	57409: {57409, []byte("OR"), "operator", 0},
	57410: {57410, []byte("and"), "operator", 0},
	57411: {57411, []byte("not"), "operator", 0},
	57398: {57398, []byte("_"), "integral", 0},
	57397: {57397, []byte("_"), "string", 0},
	57395: {57395, []byte("_"), "column name", 0},
}

// isValidToken checks if token is valid in WHERE statement
func isValidToken(i int) bool {
	_, ok := validTokens[i]
	return ok
}

func composeToken(i int, b []byte) domain.Token {
	token := validTokens[i]
	if token.TokenType != "operator" {
		token.Lexema = b
	}
	return token
}
