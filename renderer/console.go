package renderer

// Console - ..
type Console struct{}

// Table - ..
func (c Console) Table(content [][]string) string {
	output := ""
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

	output += "\n"
	return output
}
