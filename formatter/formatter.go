package formatter

// Formatter - ..
type Formatter interface {
	Parse([]map[string]string) string
	ToArray(string) []map[string]string
}
