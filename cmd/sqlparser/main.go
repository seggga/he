package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	service := services.NewService(query, &parser, checker, csv, cfg.Separator, ctx)
	go service.Run()

	sigInt := make(chan os.Signal)
	signal.Notify(sigInt, syscall.SIGINT)

	select {
	case <-ctx.Done(): // conctext has been closed due to timeout
		// log.Println("got INTERRUPT signal")
	case <-sigInt: // got INT signal
		// log.Println("got INTERRUPT signal")
		cancel()
	}

	fmt.Println("Program exit")
}

func printAppData() {
	binaryPath, err := os.Executable()
	if err != nil {
		// log.Errorf("there is a problem getting binary path, %v", err)
	}
	fmt.Printf("Path to the binary: %s\n", binaryPath)
	fmt.Printf("commit version: %s\n\n", CommitVer)
}
