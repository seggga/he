package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/seggga/he/internal/config"
)

var CommitVer string

func main() {
	binaryPath, err := os.Executable()
	// binaryName := os.Args[0]
	fmt.Printf("Path to the binary: %s\n", binaryPath)
	fmt.Printf("commit version: %s\n\n", CommitVer)

	// get application config
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Unable to read config.yaml, %v.\nProgram exit", err)
		// log.Errorf()
	}
	fmt.Printf("%+v\n", cfg)

	// obtain users query
	fmt.Print("Enter the query: ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading query, %v", err)
		return
	}
	// convert CRLF to LF
	query = strings.Replace(query, "\n", "", -1)
	// log.Debug()
	fmt.Println(query)

}
