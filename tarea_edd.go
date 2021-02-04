package main

import (
	"fmt" // manejo de informacion
	"log"
	"net/http" // manejo de servidores

	//"log" // manejo de errores en servidores
	"github.com/gorilla/mux" //usar api de manera sencilla
)

func GetDatos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo Juan Pablo Orellana 201314881")

}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", GetDatos)

	log.Fatal(http.ListenAndServe(":3000", router))

}
