package entity

import "fmt"

type BillDetail struct {
	BillDetailId string
	CustomerOrder
}

func (bd BillDetail) String() string {
	return fmt.Sprintf("Id: %s, Menu: %v, Qty: %d", bd.BillDetailId, bd.OrderedMenu, bd.Qty)
}
