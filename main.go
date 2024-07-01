package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"ascii-art-web/ascii"
)

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/web", Web)
	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Data struct {
	Res string
}

func Web(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error with the form", http.StatusMethodNotAllowed)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	words := r.Form.Get("words")
	files := r.Form.Get("Files")

	lines := strings.Split(string(words), "\n")
	result := ascii.DisplayText(files, lines)

	data := Data{
		Res: result,
	}
	err = templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
