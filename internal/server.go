package internal

import (
	"fmt"
	"net/http"
)

func StartHttpServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/create", createPage)
	http.HandleFunc("/list", listPage)
	http.ListenAndServe(":8080", nil)
}

func homePage(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	fmt.Fprint(writer, "Welcome to the homePage")
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
