package main

import (
	"gonotes/app"
	"gonotes/formatter"
	"gonotes/renderer"
	"gonotes/repository"
)

func main() {
	app := app.CliApp{
		BaseApp: app.BaseApp{
			Repository: repository.FileRepository{
				Formatter: formatter.Csv{},
				FilePath:  "./notes.csv",
			},
			Renderer: renderer.Console{},
			Schema:   []string{"content"},
		},
	}

	app.Run()
}
