package domain

type Token struct {
	Token    string
	Lexema   string
	Priority int
}

type ParsedQuery struct {
	Select []string
	Files  []string
	Where  []Token
}

type QueryReader interface {
	Read() (sql string, err error)
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

type QueryParser interface {
	Parse(string) error
	GetSelect() []string
	GetFiles() []string
	GetCondition() []Token
}

type Checker interface {
	CheckFile() error
}

type CSVDigger interface {
	SearchData(file CSVFileReader) error
}
