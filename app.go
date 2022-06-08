package main

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/usecase"
	"fmt"
)

func main() {
	fnbRepo := repository.NewFnBRepository()
	tableRepo := repository.NewTableRepository(30)
	trxRepo := repository.NewTrxRepository()

	findMenuUseCase := usecase.NewFindMenuUseCase(fnbRepo)
	f1, err := findMenuUseCase.FindMenuById("F001")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	b1, err := findMenuUseCase.FindMenuById("B001")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	name, err := findMenuUseCase.FindMenuByName("asi")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("Menu [%v", name)

	customerOrderUseCase := usecase.NewCustomerOrderUseCase(trxRepo, tableRepo)
	customerPaymentUseCase := usecase.NewCustomerPaymentUseCase(trxRepo, tableRepo)
	customer01 := entity.Customer{
		CustomerId:    "C00001",
		MobilePhoneNo: "08788123123",
		Name:          "Jution",
	}

	newBillNo, err := customerOrderUseCase.TakeOrder(customer01, "T02", []entity.CustomerOrder{
		{OrderedMenu: f1, Qty: 1},
		{OrderedMenu: b1, Qty: 2},
	})
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	//customerPaymentUseCase.OrderPayment(newBillNo)
	newBillNo, err = customerOrderUseCase.TakeOrder(customer01, "T02", []entity.CustomerOrder{
		{OrderedMenu: f1, Qty: 1},
		{OrderedMenu: b1, Qty: 2},
	})
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	customerPaymentUseCase.OrderPayment(newBillNo)
	//
	//tableViewUseCase := usecase.NewTableViewUseCase(tableRepo)
	//tableViewUseCase.ViewTable()
}
