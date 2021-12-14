package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/xwb1989/sqlparser"

	"github.com/seggga/he/internal/config"
	"github.com/seggga/he/internal/csvfile"
	"github.com/seggga/he/internal/parser"
	"github.com/seggga/he/internal/query"
	"github.com/seggga/he/internal/services"
)

var CommitVer string

func main() {
	binaryPath, err := os.Executable()
	if err != nil {
		// log.Errorf("there is a problem getting binary path, %v", err)
	}
	fmt.Printf("Path to the binary: %s\n", binaryPath)
	fmt.Printf("commit version: %s\n\n", CommitVer)

	// get application config
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Unable to read config.yaml, %v.\nProgram exit", err)
		return
	}
	fmt.Printf("%+v\n", cfg)

	query := query.NewQuery()
	parser := parser.NewParser()
	csv := new(csvfile.CSVScanner)
	service := services.NewService(query, &parser, csv)
	service.Run()

	// query, err := query.ReadQuery()
	// if err != nil {
	// 	fmt.Printf("Unable to read query, %v.\nProgram exit", err)
	// 	// log.Errorf()
	// 	return
	// }
	// // log.Debug()
	// fmt.Println(query)

	reader := strings.NewReader(`select one, two, three from file_one.csv, file_two.csv where one>10 and two="hi there"`)
	tokens := sqlparser.NewTokenizer(reader)
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			break
		}
		fmt.Println(stmt)
		// Do your logics with the statements.

	}

}
