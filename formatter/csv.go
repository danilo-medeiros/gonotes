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
	currentIndex := 0
	currentValue := ""
	ignoreComma := false

	for index, v := range raw {
		value := string(v)
		appendToCurrentValue := true
		appendToLine := false
		appendToList := false
		isLastLine := index == len(raw)-1

		if value == "\"" && string(raw[index-1]) != "\"" {
			ignoreComma = !ignoreComma
		}

		if !ignoreComma && value == "," {
			appendToLine = true
			appendToCurrentValue = false
		}

		if isLastLine {
			appendToLine = true
		}

		if value == "\n" {
			appendToCurrentValue = false
			appendToLine = true
			appendToList = !isLastLine
		}

		if appendToCurrentValue {
			currentValue += value
		}

		if appendToLine {
			valueToAppend := currentValue

			if string(currentValue[0]) == "\"" {
				valueToAppend = valueToAppend[1 : len(currentValue)-1]
			}

			valueToAppend = strings.ReplaceAll(valueToAppend, "\"\"", "\"")
			result[currentIndex] = append(result[currentIndex], valueToAppend)
			currentValue = ""
		}

		if appendToList {
			result = append(result, []string{})
			currentIndex++
		}
	}

	return result
}
