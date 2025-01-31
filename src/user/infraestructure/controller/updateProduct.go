package controller

import (
	"encoding/json"
	"net/http"
	"practica_2/src/user/application"
	"practica_2/src/user/domain"
	"practica_2/src/user/infraestructure"
	"strings"
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Validar método HTTP
	if r.Method != http.MethodPut {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Obtener ID del producto desde la URL
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 || pathParts[2] == "" {
		http.Error(w, "ID del producto requerido en la URL", http.StatusBadRequest)
		return
	}
	productID := pathParts[2] // MongoDB usa strings como _id

	// Decodificar JSON del cuerpo de la solicitud
	var updatedProduct domain.Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	// Crear instancia del repositorio
	repo := infraestructure.NewMongoDBRepository() // ✅ Ahora está bien definido
	useCase := application.NewUpdateProduct(repo) // ✅ Asegúrate de que este caso de uso exista

	// Ejecutar actualización
	err = useCase.Execute(productID, &updatedProduct)
	if err != nil {
		http.Error(w, "Error al actualizar el producto", http.StatusInternalServerError)
		return
	}

	// Respuesta de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Producto actualizado correctamente"))
}
