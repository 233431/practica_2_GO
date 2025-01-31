package application

import "practica_2/src/user/domain"


type EditProduct struct {
	repo domain.Iproduct
}

func NewupCreateProduct(repo domain.Iproduct) *EditProduct{
	return &EditProduct{repo: repo}
}

func (cp *EditProduct) Execute(id string, product *domain.Product)error{
	return cp.repo.Update(id,product)
}