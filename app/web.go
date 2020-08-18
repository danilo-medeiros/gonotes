package app

import (
	"fmt"
	"log"
	"net/http"
)

// WebApp ..
type WebApp struct {
	BaseApp
}

func (app WebApp) create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		return
	}

	value := map[string]string{
		"content": r.FormValue("content"),
	}

	result := app.Repository.Create(value)
	fmt.Fprintf(w, app.Renderer.Table([]map[string]string{result}))
}

func (app WebApp) update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		return
	}

	value := map[string]string{
		"id":      r.FormValue("id"),
		"content": r.FormValue("content"),
	}
	result := app.Repository.Update(value)
	fmt.Fprintf(w, app.Renderer.Table([]map[string]string{result}))
}

func (app WebApp) list(w http.ResponseWriter, r *http.Request) {
	list := app.Repository.ListAll()
	fmt.Fprint(w, app.Renderer.Table(list))
}

func (app WebApp) delete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		return
	}

	deleted := app.Repository.Delete(r.FormValue("id"))
	if deleted {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotModified)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status %v", http.StatusOK)
}

func (app WebApp) notes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		app.list(w, r)
	case "POST":
		app.create(w, r)
	case "PUT":
		app.update(w, r)
	case "DELETE":
		app.delete(w, r)
	default:
		app.methodNotAllowed(w, r)
	}
}

func (WebApp) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprint(w, "Method not allowed")
}

// Run - Runs the webapp and listen for the routes
func (app WebApp) Run() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/notes", app.notes)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
