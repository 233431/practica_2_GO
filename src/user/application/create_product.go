package application

import "practica_2/src/user/domain"

type ProductCreator struct {
	storage domain.Iproduct
}

func (pc *ProductCreator) Execute(product domain.Product) any {
	panic("unimplemented")
}

// Constructor para la estructura ProductCreator
func NewProductCreator(storage domain.Iproduct) *ProductCreator {
	return &ProductCreator{storage: storage}
}

// Método para registrar un nuevo producto en el repositorio
func (pc *ProductCreator) Register(product domain.Product) error {
	return pc.storage.Save(&product)
}
