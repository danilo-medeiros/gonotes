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

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status %v", http.StatusOK)
}

func (app WebApp) list(w http.ResponseWriter, r *http.Request) {
	list := app.Repository.ListAll()
	fmt.Fprint(w, app.Renderer.Table(list))
}

// Run ...
func (app WebApp) Run() {
	http.HandleFunc("/status", status)
	http.HandleFunc("/notes", app.list)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
