package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Anil", City: "Hyd", Zipcode: "500038", DOB: "1989-11-13", Status: "1"},
		{Id: "1002", Name: "Vayu", City: "Blr", Zipcode: "560084", DOB: "1989-12-13", Status: "1"},
		{Id: "1003", Name: "Vyan", City: "Chen", Zipcode: "600001", DOB: "1989-10-13", Status: "1"},
		{Id: "1004", Name: "Pavan", City: "Ral", Zipcode: "27526", DOB: "1989-09-13", Status: "1"},
	}
	return CustomerRepositoryStub{customers}
}
