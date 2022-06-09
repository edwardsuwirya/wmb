package usecase

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
)

type TableViewUseCase struct {
	tableRepo repository.TableRepository
}

func (c *TableViewUseCase) ViewTable() ([]entity.Table, error) {
	tables := c.tableRepo.FindByAvailability()
	if len(tables) == 0 {
		return nil, utils.GeneralError("Full Table")
	}
	fmt.Println(tables)
	return tables, nil
}
func (c *TableViewUseCase) GetTable(tableNo string) (entity.Table, error) {
	tables := c.tableRepo.FindById(tableNo)
	if tables.TableNo == "" {
		return entity.Table{}, utils.DataNotFoundError(tableNo)
	}
	fmt.Println(tables)
	return tables, nil
}
func NewTableViewUseCase(tableRepo repository.TableRepository) TableViewUseCase {
	return TableViewUseCase{
		tableRepo: tableRepo,
	}
}
