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
	Schema     []string
}

// Run - Runs the app
func (app App) Run() {
	fields, command := app.mapArgs(os.Args[1:])

	switch command {
	case "save":
		var result map[string]string

		if len(fields["id"]) == 0 {
			result = app.Repository.Create(fields)
		} else {
			result = app.Repository.Update(fields)
		}

		if result != nil {
			fmt.Print(app.Renderer.Table([]map[string]string{result}))
			fmt.Println("Saved note successfully")
		} else {
			fmt.Println("This note does not exist")
		}
	case "list":
		list := app.Repository.ListAll()
		fmt.Print(app.Renderer.Table(list))
		fmt.Printf("Found %v notes\n", len(list))
	case "delete":
		if app.Repository.Delete(fields["id"]) {
			fmt.Println("Deleted note successfully")
		} else {
			fmt.Println("This note does not exist")
		}
	default:
		fmt.Println("Welcome to gonotes! Insert a valid command.")
	}
}

func (app App) mapArgs(args []string) (map[string]string, string) {
	fields := make(map[string]string)
	var command string

	for index := 0; index < len(args); index++ {
		lastIndex := len(args) - 1

		if index == lastIndex {
			command = args[index]
		} else if index < lastIndex && args[index][:2] == "--" {
			field := args[index][2:]

			if contains(app.Schema, field) {
				fields[field] = args[index+1]
			}
		}
	}

	return fields, command
}

func contains(list []string, value string) bool {
	exists := false

	for _, item := range list {
		if value == item {
			exists = true
		}
	}

	return exists
}
