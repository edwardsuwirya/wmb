package repository

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
	"strings"
)

type fnbFileRepository struct {
	io *utils.FileIo
}

func (f *fnbFileRepository) readFile() []entity.FnB {
	var db []entity.FnB
	for _, d := range f.io.Read() {
		fnb := new(entity.FnB)
		utils.FromJsonString(d, fnb)
		db = append(db, *fnb)
	}
	return db
}
func (f *fnbFileRepository) FindById(id string) entity.FnB {
	db := f.readFile()
	var fnbSelected entity.FnB
	for _, fnb := range db {
		if fnb.FnBId == id {
			fnbSelected = fnb
			break
		}
	}
	return fnbSelected
}

func (f *fnbFileRepository) FindByName(name string) []entity.FnB {
	db := f.readFile()
	var fnbSelected []entity.FnB
	for _, fnb := range db {
		if strings.Contains(fnb.MenuName, name) {
			fnbSelected = append(fnbSelected, fnb)
		}
	}
	return fnbSelected
}

func (f *fnbFileRepository) initialSetup() {
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
	fnb03 := entity.FnB{
		FnBId:    "F002",
		MenuName: "Nasi Uduk",
		Price:    6000,
	}
	fnb04 := entity.FnB{
		FnBId:    "B002",
		MenuName: "Es Kopi",
		Price:    4000,
	}
	fnb05 := entity.FnB{
		FnBId:    "B003",
		MenuName: "KOpi Americano",
		Price:    7000,
	}
	db := []entity.FnB{fnb01, fnb02, fnb03, fnb04, fnb05}
	for _, menu := range db {
		jsonString := utils.ToJsonString(menu)
		f.io.Write(jsonString)
	}

	menus := f.io.Read()
	for _, menu := range menus {
		fmt.Println(menu)
	}
}

func NewFnBFileRepository(io *utils.FileIo) FnBRepository {
	repo := new(fnbFileRepository)
	repo.io = io
	repo.initialSetup()
	return repo
}
