package utils

// MapsToArrays - Creates an array of arrays (matrix) from an array of maps
func MapsToArrays(list []map[string]string) [][]string {
	result := make([][]string, 0)
	keys := make(map[string][]string)

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
		result = append(result, make([]string, len(keys)))
	}

	keyIndex := 0
	for key, values := range keys {
		result[0][keyIndex] = key

		for index, value := range values {
			result[index+1][keyIndex] = value
		}

		keyIndex++
	}

	return result
}

// ArraysToMaps - Creates an array of maps from an array of arrays
func ArraysToMaps(list [][]string) []map[string]string {
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
