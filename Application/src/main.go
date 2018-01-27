package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", runApp)
	http.HandleFunc("/home", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/static/", serveStatic)
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

func serveStatic(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, request.URL.Path[1:])
}
