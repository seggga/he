package csvfile

import (
	"os"

	"github.com/seggga/he/internal/domain"
)

type CSVFile struct {
	name    string
	sNumber int
}

func (c CSVFile) ReadHeader() ([]string, error) {
	return nil, nil
}

func (c CSVFile) NextString() (domain.DataString, error) {
	return nil, nil
}

func NewCSVFile(fileName string) (*CSVFile, error) {
	// check file existence
	if _, err := os.Stat(fileName); err != nil {
		return nil, err
	}
	return &CSVFile{
		name:    fileName,
		sNumber: 0,
	}, nil
}
