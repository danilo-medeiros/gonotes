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
			FilePath:  "./notes.csv",
		},
		Renderer: renderer.Console{},
	}

	app.Run()
}
