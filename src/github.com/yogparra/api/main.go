package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
var libros []Libro

func main() {

	router := mux.NewRouter()

	libros = append(libros,
		Libro{ID: 1, Titulo: "El señor de los anillos", Autor: "JJ tolkien", Año: "1914"},
		Libro{ID: 2, Titulo: "Golang", Autor: "Juan Gonzalez", Año: "2010"},
		Libro{ID: 3, Titulo: "Yo robot", Autor: "Asimov", Año: "1920"},
		Libro{ID: 4, Titulo: "El hobit", Autor: "JJ tolkien", Año: "1945"},
		Libro{ID: 5, Titulo: "Mi vida", Autor: "Guillermo", Año: "2016"})

	router.HandleFunc("/libros", getlibros).Methods("GET")
	router.HandleFunc("/libro/{id:[0-9]+}", getLibro).Methods("GET")
	router.HandleFunc("/libros", addLibro).Methods("POST")
	router.HandleFunc("/libros", updateLibro).Methods("PUT")
	router.HandleFunc("/libro/{id:[0-9]+}", removeLibro).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getlibros(w http.ResponseWriter, r *http.Request) {
	//log.Println("Trae todos los libros")
	json.NewEncoder(w).Encode(libros)
}

func getLibro(w http.ResponseWriter, r *http.Request) {
	//log.Println("Tra un libro por id")
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])

	for _, libro := range libros {
		if libro.ID == i {
			json.NewEncoder(w).Encode(&libro)
		}
	}

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
