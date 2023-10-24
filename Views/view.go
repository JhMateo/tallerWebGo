package Views

import (
	"fmt"
	"github.com/JhMateo/tallerWebGo/Controllers"
	"html/template" // Renderizado
	"net/http"      // Comunicación
)

func Index(w http.ResponseWriter, r *http.Request, personas []Controllers.Persona) {
	// w para escribir, r para recibir información de la petición
	template, err := template.ParseFiles("Templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "<h1>Página no encontrada!</h1>")
		panic(err)
	} else {
		template.Execute(w, personas)
	}
}

func ExecuteWeb(personas []Controllers.Persona) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Index(w, r, personas)
	})

	fsStatic := http.FileServer(http.Dir("Templates/static"))

	// Establece una ruta para servir archivos estáticos.
	http.Handle("/static/", http.StripPrefix("/static/", fsStatic))

	fmt.Println("Escuchando en http://localhost:8000...")
	http.ListenAndServe("localhost:8000", nil)
}
