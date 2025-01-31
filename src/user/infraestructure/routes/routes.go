package routes

import (
	"net/http"

	"practica_2/src/user/infraestructure/controller"
)

func LoadRoutes() {
	http.HandleFunc("/product/delete/", controller.DeleteProductHandler)
	http.HandleFunc("/products", controller.CreateProductHandler)
	http.HandleFunc("/view-products", controller.GetProductHandler)
	http.HandleFunc("/update-products/", controller.UpdateProductHandler)


}

