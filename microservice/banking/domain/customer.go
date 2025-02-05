package domain

type Customer struct {
	Id      int
	Name    string
	City    string
	Zipcode string
	DOB     string
	Staus   string
}
type ICustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(id int) (Customer, error)
	Delete(id int)
	Save(customer Customer) (Customer, error)
}
