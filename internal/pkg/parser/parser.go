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
	return p.parsedQuery.Where
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

// domain.Token: {token, lexema, tokenType, priority}
// Priority is set according to Reverse Poland Notation
var validTokens = map[int]domain.Token{
	40:    {Token: 40, Lexema: []byte("("), TokenType: "bracket", Priority: 0},
	41:    {Token: 41, Lexema: []byte(")"), TokenType: "bracket", Priority: 1},
	60:    {Token: 60, Lexema: []byte("<"), TokenType: "operator", Priority: 3},
	61:    {Token: 61, Lexema: []byte("="), TokenType: "operator", Priority: 3},
	62:    {Token: 62, Lexema: []byte(">"), TokenType: "operator", Priority: 3},
	57418: {Token: 57418, Lexema: []byte("<="), TokenType: "operator", Priority: 3},
	57419: {Token: 57419, Lexema: []byte(">="), TokenType: "operator", Priority: 3},
	57420: {Token: 57420, Lexema: []byte("!="), TokenType: "operator", Priority: 3},
	57409: {Token: 57409, Lexema: []byte("OR"), TokenType: "operator", Priority: 2},
	57410: {Token: 57410, Lexema: []byte("and"), TokenType: "operator", Priority: 2},
	57411: {Token: 57411, Lexema: []byte("not"), TokenType: "operator", Priority: 2},
	57398: {Token: 57398, Lexema: []byte("_"), TokenType: "integral", Priority: 100},
	57397: {Token: 57397, Lexema: []byte("_"), TokenType: "string", Priority: 100},
	57395: {Token: 57395, Lexema: []byte("_"), TokenType: "column name", Priority: 100},
}

// isValidToken checks if token is valid in WHERE statement
func isValidToken(i int) bool {
	_, ok := validTokens[i]
	return ok
}

func composeToken(i int, b []byte) domain.Token {
	token := validTokens[i]
	if token.TokenType != "operator" && token.TokenType != "bracket" {
		token.Lexema = b
	}
	return token
}
