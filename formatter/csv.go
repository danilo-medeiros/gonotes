package formatter

import (
	"gonotes/parser"
	"gonotes/utils"
)

// Csv - ..
type Csv struct{}

// Parse - ..
func (Csv) Parse(items []map[string]string) string {
	list := utils.MapsToArrays(items)
	parser := parser.Csv{}

	return parser.Parse(list)
}

// ToArray - ..
func (Csv) ToArray(raw string) []map[string]string {
	parser := parser.Csv{}
	result := parser.ToArray(raw)

	return utils.ArraysToMaps(result)
}
