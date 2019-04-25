package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("locaslhost:8080", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "URL.PATH = %q\n", request.URL.Path)
}
