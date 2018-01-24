package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", runApp)
	http.ListenAndServe(":8080", nil)
}

func runApp(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "<h1>Hello from Android<h1>")
}
