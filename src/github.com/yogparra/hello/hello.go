package main

import "fmt"

func main() {
	//fmt.Println(HolaMundo("Guillermo"))
	fmt.Println(HolaMundo("Andres", "español"))
}

func HolaMundo(nombre string, idioma string) string {

	var saludo = ""

	if nombre == "" {
		nombre = "Mundo"
	}
	if idioma == "español" {
		saludo = "Hola, "
	} else {
		saludo = "Salut, "
	}

	return saludo + nombre
}
