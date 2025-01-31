package application


import "practica_2/src/user/domain"

type RemoveProduct struct {
	repo domain.Iproduct
}

func NewRemoveProduct(repo domain.Iproduct) *RemoveProduct {
	return &RemoveProduct{repo: repo}
}

func (rp *RemoveProduct) Execute(id string) error {
	return rp.repo.Delete(id)
}
