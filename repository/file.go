package repository

import (
	"fmt"
	"io/ioutil"
	"time"
)

var filePath = "./notes.csv"

// FileRepository - ..
type FileRepository struct{}

// ListAll - ..
func (FileRepository) ListAll() [][]string {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil
	}

	result := make([][]string, 1)
	currentIndex := 0
	currentValue := ""

	for index, v := range data {
		value := string(v)
		appendToCurrentValue := true
		appendToLine := false
		appendToList := false
		isLastLine := index == len(data)-1

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
			currentValue += string(value)
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

// Create - ..
func (f FileRepository) Create(content string) []string {
	list := f.ListAll()
	id := time.Now().Unix()
	row := []string{fmt.Sprintf("%v", id), content}
	list = append(list, row)
	writeToFile(list)

	return row
}

// Update - ..
func (f FileRepository) Update(id string, content string) []string {
	list := f.ListAll()
	indexToUpdate := -1

	for index, value := range list {
		if value[0] == id {
			indexToUpdate = index
		}
	}

	if indexToUpdate < 0 {
		return nil
	}

	list[indexToUpdate][1] = content
	writeToFile(list)

	return list[indexToUpdate]
}

// Delete - ..
func (f FileRepository) Delete(id string) bool {
	list := f.ListAll()
	deletionIndex := -1

	for index, value := range list {
		if value[0] == id {
			deletionIndex = index
		}
	}

	if deletionIndex < 0 {
		return false
	}

	list = append(list[:deletionIndex], list[deletionIndex+1:]...)
	writeToFile(list)

	return true
}

func writeToFile(list [][]string) {
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
	ioutil.WriteFile(filePath, []byte(content), 0644)
}
