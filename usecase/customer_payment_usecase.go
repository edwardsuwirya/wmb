package usecase

import (
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
	"strings"
)

type CustomerPaymentUseCase struct {
	trxRepo   repository.TrxRepository
	tableRepo repository.TableRepository
}

func (c *CustomerPaymentUseCase) PrintBill(billNo string) {
	bill := c.trxRepo.FindById(billNo)
	var totalBill float64
	for _, billDetail := range bill.BillDetail {
		totalBill = totalBill + (float64(billDetail.Qty) * billDetail.OrderedMenu.Price)
	}
	fmt.Println(strings.Repeat("=", 100))
	fmt.Printf("Bill Items : %v\n", bill.BillDetail)
	fmt.Printf("Total for bill %s : %.2f\n", billNo, totalBill)
	fmt.Println(strings.Repeat("=", 100))
}

func (c *CustomerPaymentUseCase) OrderPayment(billNo string) error {
	bill := c.trxRepo.FindById(billNo)
	if bill.BillNo == "" {
		return utils.DataNotFoundError(billNo)
	}
	c.trxRepo.UpdateBySettled(billNo)
	c.tableRepo.UpdateAvailability(bill.TableNo.TableNo)
	fmt.Printf("Order %s successfully paid\n", billNo)
	return nil
}

func NewCustomerPaymentUseCase(trxRepo repository.TrxRepository, tableRepo repository.TableRepository) CustomerPaymentUseCase {
	return CustomerPaymentUseCase{
		trxRepo:   trxRepo,
		tableRepo: tableRepo,
	}
}
