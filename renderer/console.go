package renderer

import "gonotes/utils"

// Console - ..
type Console struct{}

// Table - ..
func (c Console) Table(list []map[string]string) string {
	output := ""
	content := utils.MapsToArrays(list)
	contentLength := make([]int, len(content[0]))

	for _, row := range content {
		for column, field := range row {
			if fieldLength := len(field); contentLength[column] < fieldLength {
				contentLength[column] = fieldLength
			}
		}
	}

	for _, row := range content {
		for column, field := range row {
			cellLength := contentLength[column] - len(field)

			output += "| "
			output += field

			for i := 0; i <= cellLength; i++ {
				output += " "
			}
		}
		output += "|\n"
	}

	return output
}
