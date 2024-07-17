package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"ascii-art-web/ascii"
)

// Data struct holds the result of the ASCII art generation
type Data struct {
	Result string
}

// mainss function calls the DisplayText function from the ascii package
// and returns the result or an error
func mainss(input []string) (string, error) {
	inputs := input[0]
	inputs = strings.ReplaceAll(inputs, "\\n", "\n")
	filenames := input[1]
	result, err := ascii.DisplayText(inputs, filenames)
	if err != nil {
		return "", err
	}
	return result, nil
}

// renderTemplate function renders the specified template with the provided data
// and sets the appropriate HTTP status code
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}, statusCode int) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	t.Execute(w, data)
}

// handler function processes the user's input and generates the ASCII art
func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		renderTemplate(w, "400.html", nil, http.StatusBadRequest)
		return
	}

	var resultBuilder strings.Builder
	if strings.Contains(r.Form.Get("Word"), "\n") {
		myslice := strings.Split(r.Form.Get("Word"), "\r\n")
		for _, word := range myslice {
			result, err := mainss([]string{word, r.Form.Get("Banner")})
			if err != nil {
				renderTemplate(w, "error.html", nil, http.StatusInternalServerError)
				return
			}
			resultBuilder.WriteString(result)
			resultBuilder.WriteString("\n")
		}
	} else {
		result, err := mainss([]string{r.Form.Get("Word"), r.Form.Get("Banner")})
		if err != nil {
			renderTemplate(w, "500.html", nil, http.StatusInternalServerError)
			return
		}
		resultBuilder.WriteString(result)
	}

	data := Data{
		Result: resultBuilder.String(),
	}

	renderTemplate(w, "index.html", data, http.StatusOK)
}

// notFoundHandler function handles the 404 Not Found error
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "404.html", nil, http.StatusNotFound)
}

// mainHandler function is the main handler for the root path "/"
// It checks if the requested URL path is not the root path, and if so,
// it calls the notFoundHandler function to handle the 404 Not Found error
// Otherwise, it calls the handler function
func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFoundHandler(w, r)
		return
	}
	handler(w, r)
}

func main() {
	// Set up the routes for the application
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/400", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "400.html", nil, http.StatusBadRequest)
	})
	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "500.html", nil, http.StatusInternalServerError)
	})
	http.HandleFunc("/404", notFoundHandler)

	// Start the server and listen on port 8081
	log.Println("starting server on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
