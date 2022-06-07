package repository

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/utils"
	"time"
)

type TrxRepository interface {
	Create(customer entity.Customer, tableNo entity.Table, orders []entity.CustomerOrder) string
	UpdateBySettled(billNo string)
	FindById(billNo string) entity.Bill
}

type trxRepository struct {
	db []entity.Bill
}

func (t *trxRepository) FindById(billNo string) entity.Bill {
	var billSelected entity.Bill
	for _, bill := range t.db {
		if bill.BillNo == billNo {
			billSelected = bill
			break
		}
	}
	return billSelected
}

func (t *trxRepository) Create(customer entity.Customer, table entity.Table, orders []entity.CustomerOrder) string {
	var billDetails []entity.BillDetail

	for _, order := range orders {
		billDetails = append(billDetails, entity.BillDetail{
			BillDetailId:  utils.GenerateId(),
			CustomerOrder: order,
		})
	}
	newBillNo := utils.GenerateId()
	newBill := entity.Bill{
		BillNo:     newBillNo,
		TableNo:    table,
		TransDate:  time.Now(),
		CustomerId: customer,
		BillDetail: billDetails,
	}
	t.db = append(t.db, newBill)
	return newBillNo
}

func (t *trxRepository) UpdateBySettled(billNo string) {
	for idx, bill := range t.db {
		if bill.BillNo == billNo {
			t.db[idx].IsSettled = true
			break
		}
	}
}

func NewTrxRepository() TrxRepository {
	repo := new(trxRepository)
	return repo
}
