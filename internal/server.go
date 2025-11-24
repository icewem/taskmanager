package internal

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title   string
	Content string
}

func StartHttpServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/create", createPage)
	http.HandleFunc("/list", listPage)
	http.ListenAndServe(":8080", nil)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	p := &Page{
		Title:   "Home Page",
		Content: "Welcome to the home page!",
	}
	t := template.Must(template.ParseFiles("internal/homePage.html"))
	err := t.Execute(writer, p)
	if err != nil {
		return
	}
}

func createPage(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/create" {
		http.NotFound(writer, request)
		return
	}
	fmt.Fprint(writer, "Welcome to the createPage")
}

func listPage(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/list" {
		http.NotFound(writer, request)
		return
	}
	fmt.Fprint(writer, "Welcome to the list")
}
