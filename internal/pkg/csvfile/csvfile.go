package csvfile

import (
	"encoding/csv"
	"os"
)

type CSVScanner struct {
	fileName string
	sNumber  int
	file     *os.File
	Reader   *csv.Reader
	head     map[string]int
	row      []string
}

func (c *CSVScanner) Init(fileName string) error {
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

// ReadHead reads the first string of the file
func (c CSVScanner) Head() ([]string, error) {
	err := c.scan()
	if err != nil {
		return nil, err
	}
	head := make(map[string]int)
	for i, v := range c.row {
		head[v] = i
	}
	c.head = head
	return c.row, err
}

// ReadRow reads a row from CSVScanner.Reader (io.Reader)
func (c CSVScanner) Row() ([]string, error) {
	err := c.scan()
	return c.row, err
}

func (c *CSVScanner) Close() {
	c.Reader = nil
	_ = c.file.Close() //nolint: errcheck // just reading file, no need to check errors
}

func (c *CSVScanner) scan() error {
	row, err := c.Reader.Read()
	c.row = row
	if err != nil {
		c.Close()
	}
	return err
}
