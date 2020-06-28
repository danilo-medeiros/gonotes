package formatter

import (
	"fmt"
	"strings"
)

// Csv - ..
type Csv struct{}

// Parse - ..
func (Csv) Parse(list [][]string) string {
	content := ""
	for _, row := range list {
		for index, field := range row {
			isLastField := index == len(row)-1
			field = strings.ReplaceAll(field, "\"", "\"\"")

			if strings.Contains(field, ",") {
				content += fmt.Sprintf("\"%s\"", field)
			} else {
				content += fmt.Sprintf("%s", field)
			}

			if isLastField {
				content += "\n"
			} else {
				content += ","
			}
		}
	}
	return content
}

// ToArray - ..
func (Csv) ToArray(raw string) [][]string {
	result := make([][]string, 1)
	lineIndex := 0
	currentValue := ""
	ignoreComma := false

	for index := 0; index < len(raw); index++ {
		value := string(raw[index])
		appendToCurrentValue := true
		appendToLine := false
		appendToList := false
		isLastChar := index == len(raw)-1

		if value == "\"" {
			if !isLastChar && string(raw[index+1]) == "\"" {
				if len(raw)-2 > index {
					index++
				}
			} else {
				ignoreComma = !ignoreComma
				appendToCurrentValue = false
			}
		}

		if value == "," && !ignoreComma {
			appendToLine = true
			appendToCurrentValue = false
		}

		if isLastChar {
			appendToLine = true
		}

		if value == "\n" {
			appendToLine = true
			appendToCurrentValue = false
			appendToList = !isLastChar
		}

		if appendToCurrentValue {
			currentValue += value
		}

		if appendToLine {
			result[lineIndex] = append(result[lineIndex], currentValue)
			currentValue = ""
		}

		if appendToList {
			result = append(result, []string{})
			lineIndex++
		}
	}

	return result
}
