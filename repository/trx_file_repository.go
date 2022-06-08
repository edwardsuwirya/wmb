package repository

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/utils"
	"time"
)

type trxFileRepository struct {
	io *utils.FileIo
}

func (t *trxFileRepository) readFile() []entity.Bill {
	var db []entity.Bill
	for _, d := range t.io.Read() {
		bill := new(entity.Bill)
		utils.FromJsonString(d, bill)
		db = append(db, *bill)
	}
	return db
}
func (t *trxFileRepository) writeFile(data []entity.Bill) {
	t.io.Clear()
	for _, data := range data {
		t.io.Write(utils.ToJsonString(data))
	}
}
func (t *trxFileRepository) FindById(billNo string) entity.Bill {
	db := t.readFile()
	var billSelected entity.Bill
	for _, bill := range db {
		if bill.BillNo == billNo {
			billSelected = bill
			break
		}
	}
	return billSelected
}

func (t *trxFileRepository) Create(customer entity.Customer, table entity.Table, orders []entity.CustomerOrder) string {
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
	t.io.Write(utils.ToJsonString(newBill))
	return newBillNo
}

func (t *trxFileRepository) UpdateBySettled(billNo string) {
	db := t.readFile()
	for idx, bill := range db {
		if bill.BillNo == billNo {
			db[idx].IsSettled = true
			break
		}
	}
	t.writeFile(db)
}

func NewTrxFileRepository(io *utils.FileIo) TrxRepository {
	repo := new(trxFileRepository)
	repo.io = io
	return repo
}
