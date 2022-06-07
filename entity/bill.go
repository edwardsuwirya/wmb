package entity

import (
	"fmt"
	"time"
)

type Bill struct {
	BillNo     string
	TableNo    Table
	TransDate  time.Time
	CustomerId Customer
	BillDetail []BillDetail
	IsSettled  bool
}

func (b Bill) String() string {
	return fmt.Sprintf(`
Bill No: %s 
Table No: [%v]
Trans Date: %v
Customer: [%v]
Bill Detail: [%v]
Settled: %v
`, b.BillNo, b.TableNo, b.TransDate, b.CustomerId, b.BillDetail, b.IsSettled)
}
