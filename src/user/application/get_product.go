package application

import "practica_2/src/user/domain"


type GetProduct struct{
	repo domain.Iproduct
}

func NewGetProduct(repo domain.Iproduct)*GetProduct{
	return &GetProduct{repo: repo}
}

func (cp *GetProduct) Execute() ([]domain.Product, error){
	return cp.repo.GetAll()
}

