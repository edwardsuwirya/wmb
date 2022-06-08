package usecase

import (
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
)

type CustomerPaymentUseCase struct {
	trxRepo   repository.TrxRepository
	tableRepo repository.TableRepository
}

func (c *CustomerPaymentUseCase) OrderPayment(billNo string) error {
	bill := c.trxRepo.FindById(billNo)
	if bill.BillNo == "" {
		return utils.BillNotFoundError{BillNo: billNo}
	} else {
		c.trxRepo.UpdateBySettled(billNo)
		c.tableRepo.UpdateAvailability(bill.TableNo.TableNo)
		fmt.Printf("Order %s successfully paid\n", billNo)
		return nil
	}
}

func NewCustomerPaymentUseCase(trxRepo repository.TrxRepository, tableRepo repository.TableRepository) CustomerPaymentUseCase {
	return CustomerPaymentUseCase{
		trxRepo:   trxRepo,
		tableRepo: tableRepo,
	}
}
