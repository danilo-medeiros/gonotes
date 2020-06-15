package renderer

import (
	"testing"
)

func getLine(line int, content string) string {
	lineContent := ""
	currentLine := 0

	for _, val := range content {
		value := string(val)

		if value == "\n" {
			if currentLine == line {
				break
			}
			lineContent = ""
			currentLine++
		} else {
			lineContent += value
		}
	}

	return lineContent
}

func TestConsole_Table_withOneLine(t *testing.T) {
	subject := Console{}
	dummyData := [][]string{
		[]string{"123456", "My first test"},
	}

	expected := "| 123456 | My first test |"
	got := getLine(0, subject.Table(dummyData))

	if got != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s\n", expected, got)
	}
}

func TestConsole_Table_withMultipleLines(t *testing.T) {
	subject := Console{}
	dummyData := [][]string{
		[]string{"123456", "First Line"},
		[]string{"654321", "Second Line"},
	}

	expected := "| 123456 | First Line  |\n| 654321 | Second Line |\n\n"
	got := subject.Table(dummyData)

	if got != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s\n", expected, got)
	}
}
