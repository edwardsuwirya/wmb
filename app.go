package main

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"fmt"
)

func main() {
	testEntity()
}

func testEntity() {
	//table01 := entity.Table{TableNo: "T01"}
	fnbRepo := repository.NewFnBRepository()
	f1 := fnbRepo.FindById("F002")
	b1 := fnbRepo.FindById("B002")
	//fmt.Println(fnbRepo.FindByName("Es"))
	//fmt.Println(fnbRepo.FindByName("asi"))

	tableRepo := repository.NewTableRepository(30)
	table02 := tableRepo.FindById("T02")
	table03 := tableRepo.FindById("T03")
	//tableRepo.UpdateAvailability("T02")
	//tableRepo.UpdateAvailability("T03")
	//fmt.Println(tableRepo.FindByAvailability())

	customer01 := entity.Customer{
		CustomerId:    "C00001",
		MobilePhoneNo: "08788123123",
		Name:          "Jution",
	}

	trxRepo := repository.NewTrxRepository()
	billNo1 := trxRepo.Create(customer01, table02, []entity.CustomerOrder{
		{OrderedMenu: f1, Qty: 1},
		{OrderedMenu: b1, Qty: 2},
	})
	fmt.Println(trxRepo.FindById(billNo1))
	billNo2 := trxRepo.Create(customer01, table03, []entity.CustomerOrder{
		{OrderedMenu: f1, Qty: 1},
	})
	fmt.Println(trxRepo.FindById(billNo2))

	trxRepo.UpdateBySettled(billNo1)
	fmt.Println(trxRepo.FindById(billNo1))
	//bill01 := entity.Bill{
	//	BillNo:     "B0000001",
	//	TableNo:    table01,
	//	TransDate:  time.Now(),
	//	CustomerId: customer01,
	//	BillDetail: []entity.BillDetail{
	//		{BillDetailId: "BD0001", OrderedMenu: fnb01, Qty: 1},
	//		{BillDetailId: "BD0002", OrderedMenu: fnb02, Qty: 1},
	//	},
	//}
	//
	//fmt.Println(bill01)
}
