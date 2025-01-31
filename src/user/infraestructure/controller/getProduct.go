package controller

import (
	"encoding/json"
	"net/http"
	"practica_2/src/user/application"
	"practice_2/src/user/infraestructure"
)



func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		repo := infraestructure.NewMongoDBRepository()
		products, err := application.NewGetProduct(repo).Execute()
		
		if err != nil {
			http.Error(w, "error al obtener los productos", http.StatusInternalServerError)
			return
		}

		response, _ := json.Marshal(products)
		w.WriteHeader(http.StatusCreated)
		w.Write(response)
	} else {
		http.Error(w, "método no+ permitido", http.StatusMethodNotAllowed)
	}
}
