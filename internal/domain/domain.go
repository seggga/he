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

type Parser interface {
	ParseSelect()
	ParseFrom()
	ParseWhere()
}

type Checker interface {
	CheckFile() error
}

type CSVDigger interface {
	SearchData(file CSVFileReader) error
}
