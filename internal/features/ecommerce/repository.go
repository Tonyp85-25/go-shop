package ecommerce

type CustomerRepository interface {
	Create(customer *Customer) (*Customer, error)
	GetByEmail(email string) (*Customer, error)
}
