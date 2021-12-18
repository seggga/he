package services

import (
	"context"
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
	ctx     context.Context
}

func NewService(q domain.QueryReader, p domain.QueryParser, checker domain.ConditionChecker, csv domain.CSVFileReader, sep string, ctx context.Context) queryService {
	return queryService{
		query:   q,
		parser:  p,
		checker: checker,
		csv:     csv,
		sep:     sep,
		ctx:     ctx,
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
		fmt.Println("error reading query, program exit")
	}
	// parse query
	err = qs.parser.Parse(sql)
	if err != nil {
		// log
		fmt.Println("error parsing the query, program exit")
	}
	qs.checker.Init(qs.parser.GetCondition())
	// if head has already been printed - don't do this again
	headNotPrinted := true
	for _, v := range qs.parser.GetFiles() {
		select {
		case <-qs.ctx.Done():
			qs.csv.Close()
			return
		default:
			// initialize csv-reader
			err := qs.csv.Init(v)
			if err != nil {
				// log
				fmt.Println("error reading file, program exit")
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
			readRow := true
			for readRow {

				select {
				case <-qs.ctx.Done():
					qs.csv.Close()
					return
				default:
					row, err := qs.csv.Row()
					if err != nil {
						if errors.Is(err, io.EOF) {
							// log debug reached the end of file
						}
						// log cannot read csv-row
						readRow = false
						break
					}
					// fmt.Printf("%v", row)
					if !readRow {
						break
					}
					// check condition
					result, err := qs.checker.Check(head, row)
					if err != nil {
						fmt.Printf("error checking condition, program exit %s", err)
						qs.csv.Close()
						return
					}
					if result {
						printData(qs.parser.GetSelect(), head, row, qs.sep)
					}
				}
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
