package controller

import (
	"net/http"
	"strings"

	"practica_2/src/user/application"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Extraemos el ID de la URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID del producto requerido", http.StatusBadRequest)
		return
	}

	id := pathParts[2]
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "ID no válido", http.StatusBadRequest)
		return
	}

	// Inyectamos el repositorio y caso de uso
	repo :=
	useCase := application.NewDeleteProduct(repo)

	// Ejecutamos la eliminación
	if err := useCase.Execute(objID); err != nil {
		http.Error(w, "Error al eliminar el producto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
