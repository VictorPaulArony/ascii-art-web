package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"ascii-art-web/ascii"
)

func mainss(input []string) string {
	inputs := input[0]
	inputs = strings.ReplaceAll(inputs, "\\n", "\n")
	filenames := input[1]
	return ascii.DisplayText(inputs, filenames)
}

type Data struct {
	Result string
}

func handler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	var resultBuilder strings.Builder
	if strings.Contains(r.Form.Get("Word"), "\n") {
		myslice := strings.Split(r.Form.Get("Word"), "\r\n")
		for _, word := range myslice {
			resultBuilder.WriteString(mainss([]string{word, r.Form.Get("Banner")}))
			resultBuilder.WriteString("\n")
		}
	} else {
		resultBuilder.WriteString(mainss([]string{r.Form.Get("Word"), r.Form.Get("Banner")}))
	}

	data := Data{
		Result: resultBuilder.String(),
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalf("ERROR PARSING TEMPLATE %v", err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatalf("ERROR EXECUTING TEMPLATE %v", err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
