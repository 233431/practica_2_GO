package controller

import (
	"encoding/json"
	"net/http"

	"practica_2/src/user/application"
	"practica_2/src/user/domain"
	"practica_2/src/user/infraestructure"

)


func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Error al procesar el JSON", http.StatusBadRequest)
		return
	}

	repo := infraestructure.NewMongoDBRepository()
	useCase := application.NewProductCreator(repo)

	if err := useCase.Execute(product); err != nil {
		http.Error(w, "Error al guardar el producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Producto creado con éxito"))
}
