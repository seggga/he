package parser

import (
	"testing"

	"github.com/seggga/he/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	testParser    = new(Parser)
	testSQL       = `select one, two, three, "where" from file_one.csv, file_two.csv where (one>10 or two="hi there") and four<=15`
	wantCondition = []domain.Token{
		{40, []byte("("), "bracket", 0},             // (
		{57395, []byte("one"), "column name", 100},  // one
		{62, []byte(">"), "operator", 3},            // >
		{57398, []byte{49, 48}, "integral", 100},    // 10
		{57409, []byte("OR"), "operator", 2},        // or
		{57395, []byte("two"), "column name", 100},  // two
		{61, []byte("="), "operator", 3},            // =
		{57397, []byte("hi there"), "string", 100},  // "hi there"
		{41, []byte(")"), "bracket", 1},             // )
		{57410, []byte("and"), "operator", 2},       // and
		{57395, []byte("four"), "column name", 100}, // four
		{57418, []byte("<="), "operator", 3},        // <=
		{57398, []byte{49, 53}, "integral", 100},    // 15
		// Token     int
		// Lexema    []byte
		// TokenType string
		// Priority  int
	}
)

func TestGetSelect(t *testing.T) {
	testParser.Parse(testSQL)

	want := []string{"one", "two", "three", "'where'"}
	assert.Equal(t, want, testParser.parsedQuery.Select, "select mismatch")
}

func TestGetFrom(t *testing.T) {
	testParser.Parse(testSQL)

	want := []string{"file_one.csv", "file_two.csv"}
	assert.Equal(t, want, testParser.parsedQuery.Files, "select mismatch")
}

func TestGetCondition(t *testing.T) {
	testParser.Parse(testSQL)
	if len(testParser.parsedQuery.Where) != len(wantCondition) {
		t.Errorf("slices differ in size: want %d, got %d", len(wantCondition), len(testParser.parsedQuery.Where))
	}
	for i, token := range wantCondition {
		if token.Token != testParser.parsedQuery.Where[i].Token {
			t.Errorf("slices differ in Token field: i %d, token want %d, got %d", i, token.Token, testParser.parsedQuery.Where[i].Token)
		}
		if string(token.Lexema) != string(testParser.parsedQuery.Where[i].Lexema) {
			t.Errorf("slices differ in Lexema field: i %d, lexema want %v, got %v", i, token.Lexema, testParser.parsedQuery.Where[i].Lexema)
		}
		if token.TokenType != testParser.parsedQuery.Where[i].TokenType {
			t.Errorf("slices differ in TokenType field: i %d, type want %s, got %s", i, token.TokenType, testParser.parsedQuery.Where[i].TokenType)
		}
		if token.Priority != testParser.parsedQuery.Where[i].Priority {
			t.Errorf("slices differ in Priority field: i %d, priority want %d, got %d", i, token.Priority, testParser.parsedQuery.Where[i].Priority)
		}
	}
}

func TestParseCondition(t *testing.T) {

	condition := `(one>10 or two="hi there") and four<=15`
	got := parseCondition(condition)
	if len(got) != len(wantCondition) {
		t.Errorf("slices differ in size: want %d, got %d", len(wantCondition), len(got))
	}
	for i, token := range wantCondition {
		if token.Token != got[i].Token {
			t.Errorf("slices differ in Token field: i %d, token want %d, got %d", i, token.Token, got[i].Token)
		}
		if string(token.Lexema) != string(got[i].Lexema) {
			t.Errorf("slices differ in Lexema field: i %d, lexema want %v, got %v", i, token.Lexema, got[i].Lexema)
		}
		if token.TokenType != got[i].TokenType {
			t.Errorf("slices differ in TokenType field: i %d, type want %s, got %s", i, token.TokenType, got[i].TokenType)
		}
		if token.Priority != got[i].Priority {
			t.Errorf("slices differ in Priority field: i %d, priority want %d, got %d", i, token.Priority, got[i].Priority)
		}
	}
}
