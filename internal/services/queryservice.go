package services

import "github.com/seggga/he/internal/domain"

type queryService struct {
	query  domain.QueryReader
	parser domain.QueryParser
}

func (qs *queryService) Run() {
	// read query
	// parse query
}

func NewService(q domain.QueryReader, p domain.QueryParser) queryService {
	return queryService{
		query:  q,
		parser: p,
	}
}
