package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testParser = new(Parser)
	testSQL    = `select one, two, three from file_one.csv, file_two.csv where (one>10 or two="hi there") and four<=10`
)

func TestGetSelect(t *testing.T) {
	testParser.Parse(testSQL)

	want := []string{"one", "two", "three"}
	assert.Equal(t, want, testParser.parsedQuery.Select, "select mismatch")
}

func TestGetFrom(t *testing.T) {
	testParser.Parse(testSQL)

	want := []string{"file_one.csv", "file_two.csv"}
	assert.Equal(t, want, testParser.parsedQuery.Files, "select mismatch")
}

func TestGetCondition(t *testing.T) {
	testParser.Parse(testSQL)

	want := `(one > 10 or two = 'hi there') and four <= 10`
	assert.Equal(t, want, testParser.condition, "select mismatch")
}
