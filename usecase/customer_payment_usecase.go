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
	var billDisc float64
	if bill.CustomerId.ActiveMember {
		billDisc = (totalBill * bill.CustomerId.DiscountPct)
	}
	fmt.Println(strings.Repeat("=", 100))
	fmt.Printf("Bill No %s\n", billNo)
	fmt.Printf("Bill Items : %v\n", bill.BillDetail)
	fmt.Printf(" Total: %21.2f\n", totalBill)
	fmt.Printf(" Disc: %21.2f\n", billDisc)
	fmt.Printf(" Grand Total: %15.2f\n", totalBill-billDisc)
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
