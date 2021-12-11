package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testSQL = `select one, two, three from file_one.csv, file_two.csv where (one>10 or two="hi there") and four<=10`
)

func TestGetSelect(t *testing.T) {
	p, _ := NewParser(testSQL)
	p.Parse()

	want := []string{"one", "two", "three"}
	assert.Equal(t, want, p.query.Select, "select mismatch")
}

func TestGetFrom(t *testing.T) {
	p, _ := NewParser(testSQL)
	p.Parse()

	want := []string{"file_one.csv", "file_two.csv"}
	assert.Equal(t, want, p.query.From, "select mismatch")
}

func TestGetCondition(t *testing.T) {
	p, _ := NewParser(testSQL)
	p.Parse()

	want := `(one > 10 or two = 'hi there') and four <= 10`
	assert.Equal(t, want, p.condition, "select mismatch")
}
