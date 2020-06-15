package renderer

// Renderer - ..
type Renderer interface {
	Table(content [][]string) string
}
