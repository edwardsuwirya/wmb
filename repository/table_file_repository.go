package repository

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
)

type tableFileRepository struct {
	io *utils.FileIo
}

func (t *tableFileRepository) readFile() []entity.Table {
	var db []entity.Table
	for _, d := range t.io.Read() {
		table := new(entity.Table)
		utils.FromJsonString(d, table)
		db = append(db, *table)
	}
	return db
}
func (t *tableFileRepository) writeFile(data []entity.Table) {
	t.io.Clear()
	for _, data := range data {
		t.io.Write(utils.ToJsonString(data))
	}
}

func (t *tableFileRepository) UpdateAvailability(id string) {
	db := t.readFile()
	for i, tbl := range db {
		if tbl.TableNo == id {
			tbl.IsAvailable = !tbl.IsAvailable
			db[i] = tbl
			break
		}
	}
	t.writeFile(db)
}

func (t *tableFileRepository) FindByAvailability() []entity.Table {
	db := t.readFile()
	var tableAvailable []entity.Table
	for _, tbl := range db {
		if tbl.IsAvailable {
			tableAvailable = append(tableAvailable, tbl)
		}
	}
	return tableAvailable
}

func (t *tableFileRepository) FindById(id string) entity.Table {
	db := t.readFile()
	var selectedTable entity.Table
	for i, tbl := range db {
		if tbl.TableNo == id {
			selectedTable = db[i]
			break
		}
	}
	return selectedTable
}

func NewTableFileRepository(io *utils.FileIo, tableCapacity int) TableRepository {
	for i := 1; i <= tableCapacity; i++ {
		newTable := entity.Table{
			TableNo:     fmt.Sprintf("T%02d", i),
			IsAvailable: true,
		}
		io.Write(utils.ToJsonString(newTable))
	}

	repo := new(tableFileRepository)
	repo.io = io
	return repo
}
