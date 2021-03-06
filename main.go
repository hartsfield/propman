package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var homeTemplate *template.Template
var inspectionTemplate *template.Template
var servicesTemplate *template.Template
var contactTemplate *template.Template

type ViewData struct {
	Name string
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("templates/home.tmpl", "templates/components/definitions.tmpl")
	if err != nil {
		panic(err)
	}
	inspectionTemplate, err = template.ParseFiles("templates/inspection.tmpl", "templates/components/definitions.tmpl")
	if err != nil {
		panic(err)
	}
	servicesTemplate, err = template.ParseFiles("templates/services.tmpl", "templates/components/definitions.tmpl")
	if err != nil {
		panic(err)
	}
	contactTemplate, err = template.ParseFiles("templates/contact.tmpl", "templates/components/definitions.tmpl")
	if err != nil {
		panic(err)
	}

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", home)
	http.HandleFunc("/inspection", inspection)
	http.HandleFunc("/services", services)
	http.HandleFunc("/contact", contact)
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

func inspection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := inspectionTemplate.Execute(w, ViewData{
		Name: "GULF",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func services(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := servicesTemplate.Execute(w, ViewData{
		Name: "GULF",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactTemplate.Execute(w, ViewData{
		Name: "GULF",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
