package services

import "github.com/seggga/he/internal/domain"

type queryService struct {
	query  domain.QueryReader
	parser domain.QueryParser
	csv    domain.CSVFileReader
}

func (qs *queryService) Run() {

	// read query
	// parse query
	// get rpn
	/* for files from query {

		init file
		defer close file

		read header
		read row

		execute rpn

		print data
	}
	*/
}

func NewService(q domain.QueryReader, p domain.QueryParser, csv domain.CSVFileReader) queryService {
	return queryService{
		query:  q,
		parser: p,
		csv:    csv,
	}
}
