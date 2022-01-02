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
	done := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	service := services.NewService(query, &parser, checker, csv, *cfg, ctx, done)
	go service.Run()

	sigInt := make(chan os.Signal, 1)
	signal.Notify(sigInt, syscall.SIGINT) //nolint: govet // syscall.SIGINT fits os.Signal interface

	select {
	case <-done: // query stopped due to timeout
		fmt.Println("finish query by timeout")
		// log.Println("got INTERRUPT signal")
	case <-sigInt: // got INT signal
		fmt.Println("finish query by TERMINATE signal")
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
