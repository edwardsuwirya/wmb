package repository

import (
	"enigmacamp.com/wmbpos/entity"
	"strings"
)

type FnBRepository interface {
	FindById(id string) entity.FnB
	FindByName(name string) []entity.FnB
}

type fnbRepository struct {
	db []entity.FnB
}

func (f *fnbRepository) FindById(id string) entity.FnB {
	var fnbSelected entity.FnB
	for _, fnb := range f.db {
		if fnb.FnBId == id {
			fnbSelected = fnb
			break
		}
	}
	return fnbSelected
}

func (f *fnbRepository) FindByName(name string) []entity.FnB {
	var fnbSelected []entity.FnB
	for _, fnb := range f.db {
		if strings.Contains(fnb.MenuName, name) {
			fnbSelected = append(fnbSelected, fnb)
		}
	}
	return fnbSelected
}

func NewFnBRepository() FnBRepository {
	repo := new(fnbRepository)
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
	repo.db = []entity.FnB{fnb01, fnb02, fnb03, fnb04}
	return repo
}
