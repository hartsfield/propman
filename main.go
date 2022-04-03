package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var homeTemplate *template.Template

type ViewData struct {
	Name string
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("templates/home.tmpl", "templates/components/definitions.tmpl")
	if err != nil {
		panic(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", home)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
		log.Printf("defaulting to port %s", port)
	}

	http.ListenAndServe(":"+port, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeTemplate.Execute(w, ViewData{
		Name: "GULF",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
