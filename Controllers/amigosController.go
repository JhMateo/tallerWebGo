package Controllers

import "fmt"

func (pe *Persona) noHayAmigos() {
	if len(pe.Amigos) == 0 {
		fmt.Println("No hay amigos.")
		return
	}
}

func (pe *Persona) ListAmigos() {
	fmt.Printf("Listando amigos de %s:\n", pe.Nombre)
	pe.noHayAmigos()

	for _, ami := range pe.Amigos {
		fmt.Println(ami, ", ")
	}
	fmt.Println()
}

func (pe *Persona) GetAmigoByID(id int) {
	fmt.Printf("Mostrando amigo de %s con el Id %d:\n", pe.Nombre, id)
	pe.noHayAmigos()

	for _, ami := range pe.Amigos {
		if ami.Id == id {
			fmt.Println(ami, "\n")
			return
		}
	}

	fmt.Println("No se encontró un amigo con el ID", id)
	fmt.Println()
}

func (pe *Persona) GetAmigoByNombre(nombre string) {
	fmt.Printf("Mostrando amigo/s de %s con el nombre %s:\n", pe.Nombre, nombre)
	pe.noHayAmigos()

	amigoEncontrado := false

	for _, ami := range pe.Amigos {
		if ami.Nombre == nombre {
			fmt.Println(ami)
			amigoEncontrado = true
		}
	}

	if !amigoEncontrado {
		fmt.Println("No se encontró un amigo con el nombre", nombre)
	}

	fmt.Println()
}

func (pe *Persona) AddAmigo(amigo Persona) {
	// Verificar que la misma persona no se agregue como amigo
	if amigo.Id == pe.Id {
		fmt.Println("No puedes agregarte a ti mismo como amigo.")
		return
	}

	// Verificar si ya es amigo
	for _, a := range pe.Amigos {
		if a.Id == amigo.Id {
			fmt.Println(amigo.Nombre, "ya es tu amigo.")
			return
		}
	}

	pe.Amigos = append(pe.Amigos, amigo)
	fmt.Println(pe.Nombre, "agregó a", amigo.Nombre, "como amigo.\n")
}

func (pe *Persona) DelAmigoById(id int) {
	fmt.Printf("Eliminando amigo de %s con el Id %d...\n\n", pe.Nombre, id)
	pe.noHayAmigos()

	for i, ami := range pe.Amigos {
		if ami.Id != id {
			return
		} else {
			pe.Amigos = append(pe.Amigos[:i], pe.Amigos[i+1:]...)
			return
		}
	}
}

func (pe *Persona) DelAmigoByNombre(nombre string) {
	fmt.Printf("Eliminando amigo/s de %s con el nombre %s...\n\n", pe.Nombre, nombre)
	pe.noHayAmigos()

	var amigosActualizados []Persona

	for _, amigo := range pe.Amigos {
		if amigo.Nombre != nombre {
			amigosActualizados = append(amigosActualizados, amigo)
		}
	}

	pe.Amigos = amigosActualizados
}
