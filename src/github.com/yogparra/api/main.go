package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Libro es la estructura
type Libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Año    string `json:"año"`
}

// Libros un array
var Libros []Libro

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/libros", getlibros).Methods("GET")
	router.HandleFunc("/libro/{id:[0-9]+}", getLibro).Methods("GET")
	router.HandleFunc("/libros", addLibro).Methods("POST")
	router.HandleFunc("/libros", updateLibro).Methods("PUT")
	router.HandleFunc("/libro/{id:[0-9]+}", removeLibro).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getlibros(w http.ResponseWriter, r *http.Request) {
	log.Println("Trae todos los libros")
}

func getLibro(w http.ResponseWriter, r *http.Request) {
	log.Println("Tra un libro por id")
}

func addLibro(w http.ResponseWriter, r *http.Request) {
	log.Println("Crea un libro")
}

func updateLibro(w http.ResponseWriter, r *http.Request) {
	log.Println("Modifica un libro")
}

func removeLibro(w http.ResponseWriter, r *http.Request) {
	log.Println("Elimina un libro")
}
