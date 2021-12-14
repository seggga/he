package rpn

import (
	"testing"

	"github.com/seggga/he/internal/domain"
)

var (
	whereTokens = []domain.Token{
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
	}

	want = []domain.Token{}
)

func TestConvertToRPN(t *testing.T) {
	want := []domain.Token{
		{57395, []byte("one"), "column name", 100},  // one
		{57398, []byte{49, 48}, "integral", 100},    // 10
		{62, []byte(">"), "operator", 3},            // >
		{57395, []byte("two"), "column name", 100},  // two
		{57397, []byte("hi there"), "string", 100},  // "hi there"
		{61, []byte("="), "operator", 3},            // =
		{57409, []byte("OR"), "operator", 2},        // or
		{57395, []byte("four"), "column name", 100}, // four
		{57398, []byte{49, 53}, "integral", 100},    // 15
		{57418, []byte("<="), "operator", 3},        // <=
		{57410, []byte("and"), "operator", 2},       // and
	}
	got := convertToRPN(whereTokens)
	if len(want) != len(got) {
		t.Errorf("slices differ in size: want %d, got %d", len(want), len(got))
	}
	for i, token := range want {
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
