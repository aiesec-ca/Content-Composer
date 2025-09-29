package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	tmplDir = "tmpl"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	tmplPath := filepath.Join(tmplDir, "editor-js-examples", tmpl)

	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Template execution error: "+err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.PathValue("file")
	renderTemplate(w, fileName+".html", nil)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/editorjs/example/{file}", indexHandler)

	fs := http.FileServer(http.Dir("./assets"))
	//mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		fs.ServeHTTP(w, r)
	})))

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
