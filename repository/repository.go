package repository

// Data - ..
type Data interface{}

// Repository - ..
type Repository interface {
	ListAll() []map[string]string
	Delete(string) bool
	Create(map[string]string) map[string]string
	Update(map[string]string) map[string]string
}
