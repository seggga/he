package csvfile

import (
	"encoding/csv"
	"strings"
	"testing"
)

var (
	testData = `col_str1,col_int,col_str2
one,1,two
three,2,four`
	testScanner = CSVScanner{
		Reader: csv.NewReader(strings.NewReader(testData)),
	}
)

func TestReadHead(t *testing.T) {
	head, _ := testScanner.ReadHead()
	want := []string{"col_str1", "col_int", "col_str2"}

	if len(want) != len(head) {
		t.Errorf("slices differ in size: want %d, got %d", len(want), len(head))
	}
	for i, v := range want {
		if head[i] != v {
			t.Errorf("slices differ in elements: i %d, want %s, got %s", i, v, head[i])
		}
	}
}

func TestReadHeadRow(t *testing.T) {
	head, _ := testScanner.ReadHead()
	want := []string{"col_str1", "col_int", "col_str2"}

	if len(want) != len(head) {
		t.Errorf("slices differ in size: want %d, got %d", len(want), len(head))
	}
	for i, v := range want {
		if head[i] != v {
			t.Errorf("slices differ in elements: i %d, want %s, got %s", i, v, head[i])
		}
	}

	row, _ := testScanner.ReadRow()
	want = []string{"one", "1", "two"}
	if len(want) != len(row) {
		t.Errorf("slices differ in size: want %d, got %d", len(want), len(head))
	}
	for i, v := range want {
		if row[i] != v {
			t.Errorf("slices differ in elements: i %d, want %s, got %s", i, v, row[i])
		}
	}

	row, _ = testScanner.ReadRow()
	want = []string{"three", "2", "four"}
	if len(want) != len(row) {
		t.Errorf("slices differ in size: want %d, got %d", len(want), len(head))
	}
	for i, v := range want {
		if row[i] != v {
			t.Errorf("slices differ in elements: i %d, want %s, got %s", i, v, row[i])
		}
	}
}
