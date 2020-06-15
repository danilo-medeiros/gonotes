package main

import (
	"gonotes/app"
	"gonotes/renderer"
	"gonotes/repository"
)

func main() {
	app := app.App{
		Repository: repository.FileRepository{},
		Renderer:   renderer.Console{},
	}

	app.Run()
}
