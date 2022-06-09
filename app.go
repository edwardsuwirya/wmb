package main

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/usecase"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
)

func main() {
	ioMenu := utils.NewFileIo("data/menu.dat")
	ioMenu.Create()
	ioMenu.Clear()
	fnbRepo := repository.NewFnBFileRepository(ioMenu)
	findMenuUseCase := usecase.NewFindMenuUseCase(fnbRepo)
	f1, err := findMenuUseCase.FindMenuById("B002")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	_, err = findMenuUseCase.FindMenuByName("asi")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	ioTable := utils.NewFileIo("data/table.dat")
	ioTable.Create()
	ioTable.Clear()
	tableRepo := repository.NewTableFileRepository(ioTable, 5)
	tableViewUseCase := usecase.NewTableViewUseCase(tableRepo)
	tableViewUseCase.ViewTable()
	//tableRepo.UpdateAvailability("T02")
	//tableViewUseCase.ViewTable()

	ioTrx := utils.NewFileIo("data/trx.dat")
	ioTrx.Create()
	trxRepo := repository.NewTrxFileRepository(ioTrx)

	ioCustomer := utils.NewFileIo("data/customer.dat")
	ioCustomer.Create()
	customerRepo := repository.NewCustomerFileRepository(ioCustomer)

	customerOrderUseCase := usecase.NewCustomerOrderUseCase(trxRepo, tableRepo, customerRepo)
	customerPaymentUseCase := usecase.NewCustomerPaymentUseCase(trxRepo, tableRepo)
	customer01 := entity.Customer{
		CustomerId:    "",
		MobilePhoneNo: "08788123123",
		Name:          "Jution",
	}
	membershipUseCase := usecase.NewMembershipUseCase(customerRepo)
	register, err := membershipUseCase.RegisterNewCustomer(customer01)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(register)

	billNo, err := customerOrderUseCase.TakeOrder(customer01, "T05", []entity.CustomerOrder{
		{OrderedMenu: f1, Qty: 1},
	})
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	customerPaymentUseCase.PrintBill(billNo)
	//err = customerPaymentUseCase.OrderPayment("8186549286549387071")
	//if err != nil {
	//	fmt.Printf("%s\n", err.Error())
	//}
	//newBillNo, err = customerOrderUseCase.TakeOrder(customer01, "T02", []entity.CustomerOrder{
	//	{OrderedMenu: f1, Qty: 1},
	//	{OrderedMenu: b1, Qty: 2},
	//})
	//if err != nil {
	//	fmt.Printf("%s\n", err.Error())
	//}
	//customerPaymentUseCase.OrderPayment(newBillNo)
	//
	//tableViewUseCase := usecase.NewTableViewUseCase(tableRepo)
	//tableViewUseCase.ViewTable()
}
