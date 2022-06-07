package main

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/usecase"
)

func main() {
	fnbRepo := repository.NewFnBRepository()
	tableRepo := repository.NewTableRepository(30)
	trxRepo := repository.NewTrxRepository()

	findMenuUseCase := usecase.NewFindMenuUseCase(fnbRepo)
	f1 := findMenuUseCase.FindMenuById("F001")
	b1 := findMenuUseCase.FindMenuById("B001")
	findMenuUseCase.FindMenuByName("asi")

	customerOrderUseCase := usecase.NewCustomerOrderUseCase(trxRepo, tableRepo)
	customerPaymentUseCase := usecase.NewCustomerPaymentUseCase(trxRepo, tableRepo)
	customer01 := entity.Customer{
		CustomerId:    "C00001",
		MobilePhoneNo: "08788123123",
		Name:          "Jution",
	}
	table02 := tableRepo.FindById("T02")
	newBillNo := customerOrderUseCase.TakeOrder(customer01, table02, []entity.CustomerOrder{
		{OrderedMenu: f1, Qty: 1},
		{OrderedMenu: b1, Qty: 2},
	})
	customerPaymentUseCase.OrderPayment(newBillNo)
	newBillNo = customerOrderUseCase.TakeOrder(customer01, table02, []entity.CustomerOrder{
		{OrderedMenu: f1, Qty: 1},
		{OrderedMenu: b1, Qty: 2},
	})
	customerPaymentUseCase.OrderPayment(newBillNo)

	tableViewUseCase := usecase.NewTableViewUseCase(tableRepo)
	tableViewUseCase.ViewTable()
}
