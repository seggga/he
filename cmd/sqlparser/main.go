package main

import (
	"fmt"
	"os"

	"github.com/seggga/he/internal/config"
	"github.com/seggga/he/internal/query"
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
		// log.Errorf()
		return
	}
	fmt.Printf("%+v\n", cfg)

	query, err := query.ReadQuery()
	if err != nil {
		fmt.Printf("Unable to read query, %v.\nProgram exit", err)
		// log.Errorf()
		return
	}
	// log.Debug()
	fmt.Println(query)

}
