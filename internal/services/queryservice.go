package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/seggga/he/internal/domain"
)

type queryService struct {
	query   domain.QueryReader
	parser  domain.QueryParser
	checker domain.ConditionChecker
	csv     domain.CSVFileReader
	cfg     domain.Config
	ctx     context.Context
	done    chan struct{}
}

func NewService(q domain.QueryReader, p domain.QueryParser, checker domain.ConditionChecker, csv domain.CSVFileReader, cfg domain.Config, ctx context.Context, done chan struct{}) queryService {
	return queryService{
		query:   q,
		parser:  p,
		checker: checker,
		csv:     csv,
		cfg:     cfg,
		ctx:     ctx,
		done:    done,
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

	runCTX, runCTXCancel := context.WithTimeout(context.Background(), time.Duration(qs.cfg.Timeout)*time.Second) //nolint: lostcancel
	defer runCTXCancel()

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
		case <-runCTX.Done():
			qs.csv.Close()
			qs.done <- struct{}{}
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
				printData(qs.parser.GetSelect(), head, head, qs.cfg.Separator)
				headNotPrinted = false
			}
			// read csv-rows
			readRow := true
			for readRow {

				select {
				case <-qs.ctx.Done():
					qs.csv.Close()
					return
				case <-runCTX.Done():
					qs.csv.Close()
					fmt.Println("finish query by timeout")
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
						printData(qs.parser.GetSelect(), head, row, qs.cfg.Separator)
					}
				}
			}
		}
	}
}

func printData(sel []string, head []string, row []string, sep string) {
	for i, v := range sel {
		for j, w := range head {
			if strings.EqualFold(v, w) {
				fmt.Printf("%s", row[j])
			}
		}
		if i < len(sel)-1 {
			fmt.Printf("%s", sep)
		}
	}
	fmt.Printf("\n")
}
