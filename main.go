package main

import (
	//"log"
	"fmt"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	fmt.Println("Servidor corriendo..")
	http.ListenAndServe(":8080", nil)

}
func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "hola Mundon")
	plantillas.ExecuteTemplate(w, "inicio", nil)

}
func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)

}
