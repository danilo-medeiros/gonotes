package formatter

// Formatter - ..
type Formatter interface {
	Parse([][]string) string
	ToArray(string) [][]string
}
