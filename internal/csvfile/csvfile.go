package csvfile

import (
	"encoding/csv"
	"os"

	"github.com/seggga/he/internal/domain"
)

type CSVScanner struct {
	fileName string
	sNumber  int
	Reader   *csv.Reader
	head     map[string]int
	row      []string
}

func (c CSVScanner) ReadHeader() ([]string, error) {
	return nil, nil
}

func (c CSVScanner) NextString() (domain.DataString, error) {
	return nil, nil
}

func (c *CSVScanner) FileInit(fileName string) error {
	// check file existence
	if _, err := os.Stat(fileName); err != nil {
		return err
	}
	c.fileName = fileName
	f, err := os.Open(fileName)
	if err != nil {
		// log.Fatal("Unable to read input file " + filePath, err)
		return err
	}

	c.Reader = csv.NewReader(f)

	return nil
}

func (c CSVScanner) FileClose() {}

func (c *CSVScanner) Scan() bool {
	a, e := c.Reader.Read()
	c.row = a
	return e == nil
}
