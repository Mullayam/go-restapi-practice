package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {

	return CustomerRepositoryStub{customers: []Customer{}}
}

// FindAll retrieves all customers from the repository
//
// Returns a slice of customer objects, or an error if the operation fails
func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {

	return c.customers, nil
}

// FindById retrieves a customer from the repository by id
//
// Returns a customer object, or an error if the operation fails
func (c CustomerRepositoryStub) FindById(id int) (Customer, error) {
	for _, customer := range c.customers {
		if customer.Id == id {
			return customer, nil
		}
	}
	return Customer{}, nil
}

// Delete deletes a customer from the repository by id
func (c CustomerRepositoryStub) Delete(id int) {

}

// Save saves a customer to the repository
//
// Returns a customer object, or an error if the operation fails
func (c CustomerRepositoryStub) Save(customer Customer) (Customer, error) {
	c.customers = append(c.customers, customer)
	return customer, nil
}
