package main

import (
	"log"
	"net/http"
)

func main() {
	index()
}
func index() {
	directorio := "./static"
	http.Handle("/", http.FileServer(http.Dir(directorio)))
	direccion := ":1995"
	log.Println("Server: " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
