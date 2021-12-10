package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testSQL = `select one, two, three from file_one.csv, file_two.csv where one>10 and two="hi there"`
)

func TestParseSelect(t *testing.T) {
	p, _ := NewParser(testSQL)
	p.ParseSelect()

	want := []string{"one", "two", "three"}
	assert.Equal(t, want, p.query.Select, "select mismatch")
}
