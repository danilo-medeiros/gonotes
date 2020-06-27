package formatter

import "fmt"

// Csv - ..
type Csv struct{}

// Parse - ..
func (Csv) Parse(list [][]string) string {
	content := ""
	for _, row := range list {
		for index, field := range row {
			isLastField := index == len(row)-1
			content += fmt.Sprintf("\"%s\"", field)

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

	for index, v := range raw {
		value := string(v)
		appendToCurrentValue := true
		appendToLine := false
		appendToList := false
		isLastLine := index == len(raw)-1

		if value == "," {
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
			result[currentIndex] = append(result[currentIndex], currentValue[1:len(currentValue)-1])
			currentValue = ""
		}

		if appendToList {
			result = append(result, []string{})
			currentIndex++
		}
	}

	return result
}
