package app

import (
	"gonotes/renderer"
	"gonotes/repository"
)

// BaseApp - Base configurations go here
type BaseApp struct {
	Repository repository.Repository
	Renderer   renderer.Renderer
	Schema     []string
}

// App - ..
type App interface {
	Run()
}
