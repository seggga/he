package query

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/seggga/he/internal/domain"
)

// NewQuery analyses users query string and produces Query structure
func NewQuery(s string) (*domain.Query, error) {
	return nil, nil
}

func ReadQuery() (string, error) {
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
