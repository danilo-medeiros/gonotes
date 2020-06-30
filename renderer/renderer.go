package renderer

// Renderer - ..
type Renderer interface {
	Table([]map[string]string) string
}
