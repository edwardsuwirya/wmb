package main

import (
	"enigmacamp.com/wmbpos/entity"
	"fmt"
	"time"
)

func main() {
	testEntity()
}

func testEntity() {
	table01 := entity.Table{TableNo: "T01"}
	fnb01 := entity.FnB{
		FnBId:    "F001",
		MenuName: "Nasi Goreng",
		Price:    15000,
	}
	fnb02 := entity.FnB{
		FnBId:    "B001",
		MenuName: "Es Teh Manis",
		Price:    4000,
	}

	customer01 := entity.Customer{
		CustomerId:    "C00001",
		MobilePhoneNo: "08788123123",
		Name:          "Jution",
	}

	bill01 := entity.Bill{
		BillNo:     "B0000001",
		TableNo:    table01,
		TransDate:  time.Now(),
		CustomerId: customer01,
		BillDetail: []entity.BillDetail{
			{BillDetailId: "BD0001", OrderedMenu: fnb01, Qty: 1},
			{BillDetailId: "BD0002", OrderedMenu: fnb02, Qty: 1},
		},
	}

	fmt.Println(bill01)
}
