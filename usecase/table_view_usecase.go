package usecase

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"fmt"
)

type TableViewUseCase struct {
	tableRepo repository.TableRepository
}

func (c *TableViewUseCase) ViewTable() []entity.Table {
	tables := c.tableRepo.FindByAvailability()
	fmt.Println(tables)
	return tables
}
func (c *TableViewUseCase) GetTable(tableNo string) entity.Table {
	tables := c.tableRepo.FindById(tableNo)
	fmt.Println(tables)
	return tables
}
func NewTableViewUseCase(tableRepo repository.TableRepository) TableViewUseCase {
	return TableViewUseCase{
		tableRepo: tableRepo,
	}
}
