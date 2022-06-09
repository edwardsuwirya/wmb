package repository

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/utils"
)

type CustomerRepository interface {
	Create(customer *entity.Customer) string
	Update(customer *entity.Customer)
	FindById(id string) entity.Customer
	FindByPhone(phoneNo string) entity.Customer
}

type customerFileRepository struct {
	io *utils.FileIo
}

func (c *customerFileRepository) readFile() []entity.Customer {
	var db []entity.Customer
	for _, d := range c.io.Read() {
		cust := new(entity.Customer)
		utils.FromJsonString(d, cust)
		db = append(db, *cust)
	}
	return db
}
func (c *customerFileRepository) writeFile(data []entity.Customer) {
	c.io.Clear()
	for _, data := range data {
		c.io.Write(utils.ToJsonString(data))
	}
}

func (c *customerFileRepository) Create(customer *entity.Customer) string {
	newId := utils.GenerateId()
	customer.CustomerId = newId
	c.io.Write(utils.ToJsonString(customer))
	return newId
}

func (c *customerFileRepository) Update(customer *entity.Customer) {
	if customer.CustomerId == "" {
		c.Create(customer)
	} else {
		db := c.readFile()
		for i, cust := range db {
			if cust.CustomerId == customer.CustomerId {
				db[i] = *customer
				break
			}
		}
		c.writeFile(db)
	}
}

func (c *customerFileRepository) FindById(id string) entity.Customer {
	db := c.readFile()
	var customerSelected entity.Customer
	for _, customer := range db {
		if customer.CustomerId == id {
			customerSelected = customer
			break
		}
	}
	return customerSelected
}

func (c *customerFileRepository) FindByPhone(phoneNo string) entity.Customer {
	db := c.readFile()
	var customerSelected entity.Customer
	for _, customer := range db {
		if customer.MobilePhoneNo == phoneNo {
			customerSelected = customer
			break
		}
	}
	return customerSelected
}

func NewCustomerFileRepository(io *utils.FileIo) CustomerRepository {
	repo := new(customerFileRepository)
	repo.io = io
	return repo
}
