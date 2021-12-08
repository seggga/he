package domain

type Query struct {
	Select []string
	From   []string
	Where  string
}

type DataString interface {
	CheckCondition() bool
	PrintString()
}

type CSVFileReader interface {
	ReadHeader() ([]string, error)
	NextString() (DataString, error)
}

type Parser struct {
	csvFiles []string
	query    Query
}

type SQLParser interface {
	CheckFile(file CSVFileReader) error
	SearchData(file CSVFileReader) error
}
