package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const templDirName = "templ"

func main() {
	mux := http.NewServeMux()

	templDir := filepath.Join("./", templDirName)

	tmpl, _ := template.ParseGlob(templDir + "**/*.html")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
