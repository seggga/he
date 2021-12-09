package domain

type Token interface {
	IsOperator() bool
}

type Query struct {
	Select []string
	From   []string
	Where  []Token
}

type DataString interface {
	CheckCondition() bool
	PrintString()
}

type CSVFileReader interface {
	Check() bool
	ReadHeader() ([]string, error)
	NextString() (DataString, error)
}

type SQLParser interface {
	CheckFile(file CSVFileReader) error
	SearchData(file CSVFileReader) error
}
