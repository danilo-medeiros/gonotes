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
func (f FileRepository) ListAll() []map[string]string {
	data, err := ioutil.ReadFile(f.FilePath)

	if err != nil {
		return nil
	}

	list := f.Formatter.ToArray(string(data))

	return list
}

// Create - ..
func (f FileRepository) Create(data map[string]string) map[string]string {
	list := f.ListAll()
	data["id"] = fmt.Sprintf("%v", time.Now().Unix())
	list = append(list, data)
	f.writeToFile(list)

	return data
}

// Update - ..
func (f FileRepository) Update(data map[string]string) map[string]string {
	list := f.ListAll()
	indexToUpdate := -1

	for index, value := range list {
		if value["id"] == data["id"] {
			indexToUpdate = index
		}
	}

	if indexToUpdate < 0 {
		return nil
	}

	list[indexToUpdate] = data
	f.writeToFile(list)

	return list[indexToUpdate]
}

// Delete - ..
func (f FileRepository) Delete(id string) bool {
	list := f.ListAll()
	deletionIndex := -1

	for index, value := range list {
		if value["id"] == id {
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

func (f FileRepository) writeToFile(list []map[string]string) {
	content := f.Formatter.Parse(list)
	ioutil.WriteFile(f.FilePath, []byte(content), 0644)
}
