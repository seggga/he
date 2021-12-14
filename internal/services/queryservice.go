package services

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/seggga/he/internal/domain"
)

type queryService struct {
	query   domain.QueryReader
	parser  domain.QueryParser
	checker domain.ConditionChecker
	csv     domain.CSVFileReader
	sep     string
}

func (qs *queryService) Run() {
	// read query
	sql, err := qs.query.Read()
	if err != nil {
		// log
		fmt.Println("program exit")
	}
	// parse query
	err = qs.parser.Parse(sql)
	if err != nil {
		// log
		fmt.Println("program exit")
	}
	// variable to check, if head has already been printed
	headNotPrinted := false
	for _, v := range qs.parser.GetFiles() {
		// initialize csv-reader
		err := qs.csv.Init(v)
		if err != nil {
			// log
			fmt.Println("Program exit")
		}
		// read head of the csv-file
		head, err := qs.csv.Head()
		if err != nil {
			// log cannot read csv-head
			continue
		}
		if headNotPrinted {
			printData(qs.parser.GetSelect(), head, head, qs.sep)
			headNotPrinted = true
		}
		// read csv-rows
		for {
			row, err := qs.csv.Row()
			if err != nil {
				if errors.Is(err, io.EOF) {
					// log debug reached the end of file
					break
				}
				// log cannot read csv-row
				break
			}
			// check condition
			if qs.checker.Check(head, row) {
				printData(qs.parser.GetSelect(), head, row, qs.sep)
			}
		}

	}
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

func NewService(q domain.QueryReader, p domain.QueryParser, checker domain.ConditionChecker, csv domain.CSVFileReader, sep string) queryService {
	return queryService{
		query:   q,
		parser:  p,
		checker: checker,
		csv:     csv,
		sep:     sep,
	}
}

func printData(sel []string, head []string, row []string, sep string) {
	for i, v := range sel {
		for j, w := range head {
			if strings.ToLower(v) == strings.ToLower(w) {
				fmt.Printf("%s", row[j])
			}
		}
		if i < len(sel)-1 {
			fmt.Print("%s", sep)
		}
	}
	fmt.Printf("\n")
}
