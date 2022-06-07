package usecase

import (
	"enigmacamp.com/wmbpos/repository"
	"fmt"
)

type CustomerPaymentUseCase struct {
	trxRepo   repository.TrxRepository
	tableRepo repository.TableRepository
}

func (c *CustomerPaymentUseCase) OrderPayment(billNo string) {
	bill := c.trxRepo.FindById(billNo)
	bill.IsSettled = true
	c.tableRepo.UpdateAvailability(bill.TableNo.TableNo)
	fmt.Printf("Order %s successfully paid\n", billNo)
}

func NewCustomerPaymentUseCase(trxRepo repository.TrxRepository, tableRepo repository.TableRepository) CustomerPaymentUseCase {
	return CustomerPaymentUseCase{
		trxRepo:   trxRepo,
		tableRepo: tableRepo,
	}
}
