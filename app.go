package main

import (
	"enigmacamp.com/wmbpos/repository"
	"fmt"
)

func main() {
	testEntity()
}

func testEntity() {
	//table01 := entity.Table{TableNo: "T01"}
	fnbRepo := repository.NewFnBRepository()
	fmt.Println(fnbRepo.FindById("F002"))
	fmt.Println(fnbRepo.FindByName("Es"))
	fmt.Println(fnbRepo.FindByName("asi"))

	//customer01 := entity.Customer{
	//	CustomerId:    "C00001",
	//	MobilePhoneNo: "08788123123",
	//	Name:          "Jution",
	//}
	//
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
