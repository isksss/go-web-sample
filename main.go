package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

var templates = make(map[string]*template.Template)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("public/"))))
	port := "8000"
	templates["index"] = loadTemplate("index")
	http.HandleFunc("/", handleIndex)
	log.Printf("Server listening on port %s", port)
	log.Print(http.ListenAndServe(":"+port, nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := templates["index"].Execute(w, struct {
		Header  header
		Message string
		Time    time.Time
	}{
		Header:  newHeader("test"),
		Message: "hello world.",
		Time:    time.Now(),
	}); err != nil {
		log.Printf("failed to execute template: %v", err)
	}
}

func loadTemplate(name string) *template.Template {
	t, err := template.ParseFiles(
		"template/"+name+".html",
		"template/_header.html",
		"template/_footer.html",
	)
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	return t
}

type header struct {
	Title    string
	UserName string
}

func newHeader(title string) header {
	return header{Title: title, UserName: "Bob"}
}
