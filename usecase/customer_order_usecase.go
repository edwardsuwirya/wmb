package usecase

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
)

type CustomerOrderUseCase struct {
	trxRepo      repository.TrxRepository
	tableRepo    repository.TableRepository
	customerRepo repository.CustomerRepository
}

func (c *CustomerOrderUseCase) TakeOrder(customer entity.Customer, tableNo string, orders []entity.CustomerOrder) (string, error) {
	var newBillNo string
	if customer.MobilePhoneNo == "" {
		return "", utils.RequiredError("Customer Phone")
	}
	cust := c.customerRepo.FindByPhone(customer.MobilePhoneNo)

	// Auto Customer Registration
	if cust.MobilePhoneNo == "" {
		c.customerRepo.Create(&customer)
	}
	tableReserve := c.tableRepo.FindById(tableNo)
	fmt.Println("==", tableReserve)
	if tableReserve.IsAvailable {
		newBillNo = c.trxRepo.Create(cust, tableReserve, orders)
		c.tableRepo.UpdateAvailability(tableNo)
		fmt.Printf("Order %s successfully created\n", newBillNo)
	} else {
		return "", utils.TableUnavailableError(tableNo)
	}
	return newBillNo, nil
}

func NewCustomerOrderUseCase(trxRepo repository.TrxRepository, tableRepo repository.TableRepository, customerRepo repository.CustomerRepository) CustomerOrderUseCase {
	return CustomerOrderUseCase{
		trxRepo:      trxRepo,
		tableRepo:    tableRepo,
		customerRepo: customerRepo,
	}
}
