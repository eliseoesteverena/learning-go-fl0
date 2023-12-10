package main

import (
	"html/template"
	"log"
	"net/http"
)

//Estructuras
type Usuario struct {
	UserName string
	Edad     int
}

//Handlers
func Index(rw http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("index.html")

	usuario := Usuario{"Eliseo", 28}

	if err != nil {
		panic(err)
	} else {
		template.Execute(rw, usuario)
	}
}

//Función principal
func main() {

	//Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	//Server
	server := &http.Server{
		Addr:    "localhost:1995",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
