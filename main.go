package main

import (
	// Imprimir en consola

	"encoding/json"
	"fmt" // Ayuda a escribir en la respuesta
	"io"
	"log"      //Loguear si algo sale mal
	"net/http" // El paquete HTTP
)

type Data struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Cp      string `json:"cp"`
	City    string `json:"city"`
	State   string `json:"state"`
}
type GroupData struct {
	Remitente    Data `json:"remitente"`
	Destinatario Data `json:"destinatario"`
}

func main() {
	index()
}
func index() {
	http.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Se espera un método POST", http.StatusMethodNotAllowed)
			return
		}

		// Decodificar el cuerpo JSON recibido en la estructura definida
		var datosRecibidos GroupData
		err := json.NewDecoder(r.Body).Decode(&datosRecibidos)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Convertir datos en datos JSON
		datosJson, err := json.Marshal(datosRecibidos)
		if err != nil {
			log.Fatal(err)
		}
		// Decodificar JSON en una estructura
		datoStruct := GroupData{}
		error := json.Unmarshal([]byte(datosJson), &datoStruct)
		if error != nil {
			log.Fatal(error)
		}
		fmt.Printf("%+v\n", datoStruct.Remitente.Address)

		io.WriteString(w, string(datoStruct.Remitente.Address))

		respuestaJson, err := json.Marshal(datoStruct)
		if err != nil {
			log.Fatal(err)
		}

		_, _ = w.Write(respuestaJson)
	})

	directorio := "./static"
	http.Handle("/", http.FileServer(http.Dir(directorio)))

	direccion := ":1995"
	log.Println("Server: " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
