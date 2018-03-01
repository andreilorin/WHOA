package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	address = flag.String("addr", ":8080", "http service address")
	data    map[string]string
)

func main() {
	flag.Parse()
	data = map[string]string{}

	router := httprouter.New()

	router.GET("/entry/:key", show)
	router.GET("/list", show)
	router.PUT("/entry/:key/:value", update)
	router.DELETE("/delete/:key", remove)

	http.HandleFunc("/", runApp)
	http.HandleFunc("/home", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/insurance/", serveInsuranceMain)

	log.Fatal(http.ListenAndServe(*address, router))
}

func show(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	key := params.ByName("key")
	if key == "" {
		fmt.Fprintf(writer, "Read list: %v", data)
		return
	}
	fmt.Fprintf(writer, "Read entry: data[%s]= %s", key, data[key])
}

func update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	key := params.ByName("key")
	value := params.ByName("value")

	data[key] = value

	fmt.Fprintf(writer, "Updated: data[%s] = %s", key, data[key])
}

func remove(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	key := params.ByName("key")

	delete(data, key)

	fmt.Fprintf(writer, "Removed: data[%s]", key)
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

func serveInsuranceMain(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, request.URL.Path[1:])
}
