package main

import (
	"gonotes/app"
	"gonotes/formatter"
	"gonotes/renderer"
	"gonotes/repository"
)

func main() {
	app := app.App{
		Repository: repository.FileRepository{
			Formatter: formatter.Csv{},
		},
		Renderer: renderer.Console{},
	}

	app.Run()
}
