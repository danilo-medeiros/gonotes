package app

import (
	"fmt"
	"gonotes/renderer"
	"gonotes/repository"
	"os"
)

// App - ..
type App struct {
	Repository repository.Repository
	Renderer   renderer.Renderer
}

// Run - Runs the app
func (app App) Run() {
	var id string
	var command string
	var content string
	args := os.Args[1:]

	for i, v := range args {
		if v == "--id" {
			id = args[i+1]
			continue
		}

		if v == "--content" {
			content = args[i+1]
			continue
		}
	}

	command = args[len(args)-1]

	switch command {
	case "save":
		var result []string

		if len(id) == 0 {
			result = app.Repository.Create(content)
		} else {
			result = app.Repository.Update(id, content)
		}

		if result != nil {
			fmt.Print(app.Renderer.Table([][]string{result}))
			fmt.Println("Saved note successfully")
		} else {
			fmt.Println("This note does not exist")
		}
	case "list":
		list := app.Repository.ListAll()
		fmt.Print(app.Renderer.Table(list))
		fmt.Printf("Found %v notes\n", len(list))
	case "delete":
		if app.Repository.Delete(id) {
			fmt.Println("Deleted note successfully")
		} else {
			fmt.Println("This note does not exist")
		}
	default:
		fmt.Println("Welcome to gonotes! Insert a valid command.")
	}
}
