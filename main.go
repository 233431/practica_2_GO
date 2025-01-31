package main

import (
	"fmt"
	"practica_2/src/core" // Asegúrate de que este sea el path correcto
)

func main() {
	// Llamar a la función de conexión
	core.ConnectDB()

	// Obtener una colección y usarla
	collection := core.GetCollection("usuarios")
	fmt.Println("Colección obtenida:", collection)
}



