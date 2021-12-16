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

func NewService(q domain.QueryReader, p domain.QueryParser, checker domain.ConditionChecker, csv domain.CSVFileReader, sep string) queryService {
	return queryService{
		query:   q,
		parser:  p,
		checker: checker,
		csv:     csv,
		sep:     sep,
	}
}

// Run executes:
//    reads query from console,
//    parses query,
//    reads csv-rows
//    decides whether to print them or not
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
	// if head has already been printed - don't do this again
	headNotPrinted := true
	for _, v := range qs.parser.GetFiles() {
		// initialize csv-reader
		err := qs.csv.Init(v)
		if err != nil {
			// log
			fmt.Println("Program exit")
			return
		}
		// read head of the csv-file
		head, err := qs.csv.Head()
		if err != nil {
			// log cannot read csv-head
			continue
		}
		// if head has already been printed - don't do this again
		if headNotPrinted {
			printData(qs.parser.GetSelect(), head, head, qs.sep)
			headNotPrinted = false
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
}

func printData(sel []string, head []string, row []string, sep string) {
	for i, v := range sel {
		for j, w := range head {
			if strings.ToLower(v) == strings.ToLower(w) {
				fmt.Printf("%s", row[j])
			}
		}
		if i < len(sel)-1 {
			fmt.Printf("%s", sep)
		}
	}
	fmt.Printf("\n")
}
