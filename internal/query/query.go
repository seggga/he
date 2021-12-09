package query

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// NewQuery analyses users query string and produces Query structure
func NewQuery() (string, error) {
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
	return query, nil
}
