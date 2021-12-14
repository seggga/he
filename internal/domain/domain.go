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

// type DataString interface {
// 	CheckCondition() bool
// 	PrintString()
// }

type CSVFileReader interface {
	ReaderInit(string) error
	ReadHead() ([]string, error)
	ReadRow() ([]string, error)
}

type QueryParser interface {
	Parse(string) error
	GetSelect() []string
	GetFiles() []string
	GetCondition() []Token
}

type ConditionChecker interface {
	Check() bool
}
