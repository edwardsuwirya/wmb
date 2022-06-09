package usecase

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
)

type CustomerOrderUseCase struct {
	trxRepo   repository.TrxRepository
	tableRepo repository.TableRepository
}

func (c *CustomerOrderUseCase) TakeOrder(customer entity.Customer, tableNo string, orders []entity.CustomerOrder) (string, error) {
	var newBillNo string
	tableReserve := c.tableRepo.FindById(tableNo)
	if tableReserve.IsAvailable {
		newBillNo = c.trxRepo.Create(customer, tableReserve, orders)
		c.tableRepo.UpdateAvailability(tableNo)
		fmt.Printf("Order %s successfully created\n", newBillNo)
	} else {
		return "", utils.TableUnavailableError(tableNo)
	}
	return newBillNo, nil
}

func NewCustomerOrderUseCase(trxRepo repository.TrxRepository, tableRepo repository.TableRepository) CustomerOrderUseCase {
	return CustomerOrderUseCase{
		trxRepo:   trxRepo,
		tableRepo: tableRepo,
	}
}
