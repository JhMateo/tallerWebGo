package main

import (
	"github.com/JhMateo/tallerWebGo/Controllers"
	"github.com/JhMateo/tallerWebGo/Views"
)

func main() {
	per0 := Controllers.Persona{Id: 1, Nombre: "Maria", Apellido: "Gomez", Edad: 30}
	per1 := Controllers.Persona{Id: 2, Nombre: "Pepe", Apellido: "Rodriguez", Edad: 29}
	per2 := Controllers.Persona{Id: 3, Nombre: "Maria", Apellido: "Magdalena", Edad: 25}

	per0.Saludar()

	Controllers.ListarPersonas(per0, per1, per2)

	per1.AddAmigo(per0)
	per1.AddAmigo(per2)
	per1.AddAmigo(Controllers.Persona{Id: 4, Nombre: "Enrique", Apellido: "Perez", Edad: 56})

	per1.GetAmigoByNombre("Maria")
	per1.GetAmigoByID(3)

	per1.DelAmigoByNombre("Maria")
	per1.GetAmigoByNombre("Maria")

	per2.Actualizar("Josefina", "Magdalena", 25)

	per1.AddAmigo(per2)
	per1.ListAmigos()
	per1.DelAmigoById(3)
	per1.ListAmigos()
	per0.ListAmigos()
	per2.ListAmigos()

	per3 := Controllers.Persona{Id: 5, Nombre: "Juanito", Apellido: "Alimaña", Edad: 19}
	per4 := Controllers.Persona{Id: 6, Nombre: "Gabriel", Apellido: "Marquez", Edad: 50}
	per5 := Controllers.Persona{Id: 7, Nombre: "Luisito", Apellido: "Comunica", Edad: 25}

	per0.AddAmigo(per1)
	per0.AddAmigo(per3)
	per0.AddAmigo(per4)
	per0.AddAmigo(per5)

	per2.AddAmigo(per0)
	per2.AddAmigo(per1)
	per2.AddAmigo(per5)

	per4.AddAmigo(per5)

	per5.AddAmigo(per4)

	personas := Controllers.ListarPersonas(per0, per1, per2, per3, per4, per5)

	Views.ExecuteWeb(personas)
}
