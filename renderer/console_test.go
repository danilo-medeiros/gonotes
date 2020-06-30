package renderer

import (
	"testing"
)

func TestConsole_Table_withOneLine(t *testing.T) {
	subject := Console{}
	dummyData := []map[string]string{
		{
			"id":      "123456",
			"content": "My first test",
		},
	}

	expected := "| id     | content       |\n| 123456 | My first test |\n"
	got := subject.Table(dummyData)

	if got != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s\n", expected, got)
	}
}

func TestConsole_Table_withMultipleLines(t *testing.T) {
	subject := Console{}
	dummyData := []map[string]string{
		{
			"id":      "123456",
			"content": "First Line",
		},
		{
			"id":      "654321",
			"content": "Second Line",
		},
	}

	expected := "| id     | content     |\n| 123456 | First Line  |\n| 654321 | Second Line |\n"
	got := subject.Table(dummyData)

	if got != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s\n", expected, got)
	}
}
