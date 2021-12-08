package csvdata

import "fmt"

type CSVdataString struct {
	header []string // TODO: проанализиовать необходимость в этом поле
	blob   map[string]string
}

func (ds CSVdataString) CheckCondition() bool {
	return true
}

func (ds CSVdataString) PrintString() {
	fmt.Println("here is a string")
}

func NewDataString(header []string, data []string) CSVdataString {
	blob := make(map[string]string, len(header))
	for i, v := range header {
		blob[v] = data[i]
	}
	return CSVdataString{
		header: header,
		blob:   blob,
	}
}
