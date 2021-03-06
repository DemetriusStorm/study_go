package main

import (
	"fmt"
	"net/http"

	"github.com/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlsshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":      "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//Build the YAMLHandler using the mapHandler as the fallback
	yaml := `
	- path: /urlshort
	  url: https://github.com/gophercises/urlshort
	- path: /urlsgort-final
	  url: https://github.com/gophercises/urlshort/tree/final
	`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the serbver on :8080")
	http.ListenAndServe(":8080", yamlHandler)
	//http.ListenAndServe(":8080", mapHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello, world!")
}
