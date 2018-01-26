package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", runApp)
	http.HandleFunc("/home", home)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func runApp(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "<h1>App is running<h1>")
}

func home(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "<h1>Homepage<h1>")
}

func about(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "<h1>About page<h1>")
}
