package Controllers

import (
	"fmt"
)

var Cons int = 33333

type Persona struct {
	Id               int
	Nombre, Apellido string
	Edad             int
	Amigos           []Persona
}

func (pe *Persona) Saludar() {
	fmt.Println("Hola soy ", pe.Nombre, " ",
		pe.Apellido, " y tengo ", pe.Edad, " años de edad.\n")
}

func ListarPersonas(personas ...Persona) []Persona {
	fmt.Println("Listado de Personas:")
	for _, persona := range personas {
		fmt.Printf("%+v\n", persona)
	}
	fmt.Println()

	return personas
}

func (pe *Persona) Actualizar(nombre, apellido string, edad int) {
	pe.Nombre = nombre
	pe.Apellido = apellido
	pe.Edad = edad
	fmt.Println("Datos actualizados para", pe.Nombre, pe.Apellido)
	fmt.Println()
}
