package controller

import (
	"encoding/json"
	"net/http"
	"practica_2/src/user/application"
	"practica_2/src/user/infraestructure"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Creación de la instancia del repositorio MongoDB
		repo := infraestructure.NewMongoDBRepository()

		// Creación del caso de uso 'GetProduct' con el repositorio
		useCase := application.NewGetProduct(repo)

		// Ejecutar el caso de uso
		products, err := useCase.Execute()
		
		if err != nil {
			http.Error(w, "error al obtener los productos", http.StatusInternalServerError)
			return
		}

		// Serializar los productos a formato JSON
		response, _ := json.Marshal(products)

		// Enviar respuesta con los productos
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
	}
}
