package main

import (
	"fmt"
	"os"

	"github.com/seggga/he/internal/pkg/config"
	"github.com/seggga/he/internal/pkg/csvfile"
	"github.com/seggga/he/internal/pkg/parser"
	"github.com/seggga/he/internal/pkg/query"
	"github.com/seggga/he/internal/pkg/rpn"
	"github.com/seggga/he/internal/services"
)

var CommitVer string

func main() {
	printAppData()
	// get application config
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Unable to read config.yaml, %v.\nProgram exit", err)
		return
	}

	query := query.NewQuery()
	parser := parser.NewParser()
	checker := new(rpn.ConditionCheck)
	csv := new(csvfile.CSVScanner)

	service := services.NewService(query, &parser, checker, csv, cfg.Separator)
	service.Run()

}

func printAppData() {
	binaryPath, err := os.Executable()
	if err != nil {
		// log.Errorf("there is a problem getting binary path, %v", err)
	}
	fmt.Printf("Path to the binary: %s\n", binaryPath)
	fmt.Printf("commit version: %s\n\n", CommitVer)
}
