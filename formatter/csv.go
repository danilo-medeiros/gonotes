package formatter

import (
	"gonotes/parser"
)

// Csv - ..
type Csv struct{}

// Parse - ..
func (Csv) Parse(items []map[string]string) string {
	list := mapsToArrays(items)
	parser := parser.Csv{}

	return parser.Parse(list)
}

// ToArray - ..
func (Csv) ToArray(raw string) []map[string]string {
	parser := parser.Csv{}
	result := parser.ToArray(raw)

	return arraysToMaps(result)
}

func mapsToArrays(list []map[string]string) [][]string {
	result := make([][]string, 0)
	keys := make(map[string][]string, 0)

	for _, row := range list {
		for key, value := range row {
			values, exists := keys[key]
			if !exists {
				keys[key] = []string{value}
			} else {
				keys[key] = append(values, value)
			}
		}
	}

	for index := 0; index <= len(list); index++ {
		result[0] = make([]string, 0)
	}

	keyIndex := 0
	for key, values := range keys {
		result[0] = append(result[0], key)

		for index, value := range values {
			result[index+1][keyIndex] = value
		}
		keyIndex++
	}

	return result
}

func arraysToMaps(list [][]string) []map[string]string {
	keys := list[0]
	result := make([]map[string]string, 0)

	for i := 1; i < len(list); i++ {
		row := make(map[string]string)

		for j, key := range keys {
			row[key] = list[i][j]
		}

		result = append(result, row)
	}

	return result
}
