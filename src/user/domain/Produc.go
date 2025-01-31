package domain

type Iproduct interface {
	Save(product *Product)error
	GetAll()([]Product,error)
	Delete(id string)error
	Update(id int,product *Product)error
}