package repository

// Data - ..
type Data interface{}

// Repository - ..
type Repository interface {
	ListAll() [][]string
	Delete(id string) bool
	Create(content string) []string
	Update(id string, content string) []string
}
