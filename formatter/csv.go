package formatter

import (
	"encoding/csv"
	"gonotes/utils"
	"log"
	"strings"
)

// Csv - ..
type Csv struct{}

// Parse - ..
func (Csv) Parse(items []map[string]string) string {
	list := utils.MapsToArrays(items)
	sb := strings.Builder{}
	w := csv.NewWriter(&sb)
	err := w.WriteAll(list)

	if err != nil {
		log.Fatal(err)
	}

	return sb.String()
}

// ToArray - ..
func (Csv) ToArray(raw string) []map[string]string {
	r := csv.NewReader(strings.NewReader(raw))
	result, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return utils.ArraysToMaps(result)
}
