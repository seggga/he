package services

import (
	"github.com/seggga/he/internal/domain"
)

type Parser struct {
	QueryString string
	Query       domain.Query
}

func (p *Parser) Run() {
	// read query
	// parse query
}

func NewParser(queryString string) Parser {
	return Parser{
		QueryString: queryString,
	}
}
