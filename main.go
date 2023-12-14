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
	http.HandleFunc("/info", sentInfo)
	direccion := ":1995"
	log.Println("Server: " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
func sentInfo(w http.ResponseWriter, r *http.Request) {
	str := "<h1 font-color=\"Red\">Hola Jhon!</h1>"
	w.Write([]byte(str))
}
