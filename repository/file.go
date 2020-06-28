package repository

import (
	"fmt"
	"gonotes/formatter"
	"io/ioutil"
	"time"
)

// FileRepository - ..
type FileRepository struct {
	Formatter formatter.Formatter
	FilePath  string
}

// ListAll - ..
func (f FileRepository) ListAll() [][]string {
	data, err := ioutil.ReadFile(f.FilePath)

	if err != nil {
		return nil
	}

	parsedContent := string(data)
	return f.Formatter.ToArray(parsedContent)
}

// Create - ..
func (f FileRepository) Create(content string) []string {
	list := f.ListAll()
	id := time.Now().Unix()
	row := []string{fmt.Sprintf("%v", id), content}
	list = append(list, row)
	f.writeToFile(list)

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
	f.writeToFile(list)

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
	f.writeToFile(list)

	return true
}

func (f FileRepository) writeToFile(list [][]string) {
	content := f.Formatter.Parse(list)
	ioutil.WriteFile(f.FilePath, []byte(content), 0644)
}
