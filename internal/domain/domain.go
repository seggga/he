package domain

type Token struct {
	Token     int
	Lexema    []byte
	TokenType string
	Priority  int
}

type ParsedQuery struct {
	Select   []string
	Files    []string
	Where    []Token
	ColNames []string
}

type QueryReader interface {
	Read() (sql string, err error)
}

type DataString interface {
	CheckCondition() bool
	PrintString()
}

type CSVFileReader interface {
	FileInit(string) error
	FileClose()
	ReadHeader() ([]string, error)
	ReadRow() (DataString, error)
}

type QueryParser interface {
	Parse(string) error
	GetSelect() []string
	GetFiles() []string
	GetCondition() []Token
}
