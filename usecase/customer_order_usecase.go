package usecase

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"fmt"
)

type CustomerOrderUseCase struct {
	trxRepo   repository.TrxRepository
	tableRepo repository.TableRepository
}

func (c *CustomerOrderUseCase) TakeOrder(customer entity.Customer, tableNo string, orders []entity.CustomerOrder) string {
	var newBillNo string
	tableReserve := c.tableRepo.FindById(tableNo)
	if tableReserve.IsAvailable {
		newBillNo = c.trxRepo.Create(customer, tableReserve, orders)
		c.tableRepo.UpdateAvailability(tableReserve.TableNo)
		fmt.Printf("Order %s successfully created\n", newBillNo)
	} else {
		fmt.Printf("Table %s is not available\n", tableReserve.TableNo)
	}
	return newBillNo
}

func NewCustomerOrderUseCase(trxRepo repository.TrxRepository, tableRepo repository.TableRepository) CustomerOrderUseCase {
	return CustomerOrderUseCase{
		trxRepo:   trxRepo,
		tableRepo: tableRepo,
	}
}
