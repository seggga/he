package query

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Query struct{}

// NewQuery - ...
func NewQuery() Query {
	return Query{}
}

// NewQuery analyzes users query string and produces Query structure
func (q Query) Read() (string, error) {
	return readQuery()
}

func readQuery() (string, error) {
	// obtain users query
	fmt.Print("Enter the query: ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading query, %v", err)
		return "", err
	}
	// convert CRLF to LF
	query = strings.Replace(query, "\n", "", -1)
	// query := `select col1, col2, col3 from pat.csv, mat.csv where col1>1 and col3="bar"`
	return query, nil
}
