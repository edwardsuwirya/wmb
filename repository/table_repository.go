package repository

import (
	"enigmacamp.com/wmbpos/entity"
	"fmt"
)

type TableRepository interface {
	UpdateAvailability(id string)
	FindByAvailability() []entity.Table
}

type tableRepository struct {
	db []entity.Table
}

func (t *tableRepository) UpdateAvailability(id string) {
	//For loop pass By Value, tricky
	//for _, tbl := range t.db {
	//	if tbl.TableNo == id {
	//		tbl.IsAvailable = !tbl.IsAvailable
	//		break
	//	}
	//}
	for i, tbl := range t.db {
		if tbl.TableNo == id {
			tbl.IsAvailable = !tbl.IsAvailable
			t.db[i] = tbl
			break
		}
	}
}

func (t *tableRepository) FindByAvailability() []entity.Table {
	var tableAvailable []entity.Table
	for _, tbl := range t.db {
		if tbl.IsAvailable {
			tableAvailable = append(tableAvailable, tbl)
		}
	}
	return tableAvailable
}

func NewTableRepository(tableCapacity int) TableRepository {
	tableDb := make([]entity.Table, tableCapacity)
	for i := 1; i <= tableCapacity; i++ {
		newTable := entity.Table{
			TableNo:     fmt.Sprintf("T%02d", i),
			IsAvailable: true,
		}
		tableDb[i-1] = newTable
	}
	repo := new(tableRepository)
	repo.db = tableDb
	return repo
}
