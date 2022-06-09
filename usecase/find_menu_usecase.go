package usecase

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
)

type FindMenuUseCase struct {
	fnbRepo repository.FnBRepository
}

func (c *FindMenuUseCase) FindMenuById(id string) (entity.FnB, error) {
	fnb := c.fnbRepo.FindById(id)
	if fnb.FnBId == "" {
		return entity.FnB{}, utils.DataNotFoundError(id)
	} else {
		fmt.Printf("Product found : [%v]\n", fnb)
	}
	return fnb, nil
}
func (c *FindMenuUseCase) FindMenuByName(keyword string) ([]entity.FnB, error) {
	fnbs := c.fnbRepo.FindByName(keyword)
	if len(fnbs) == 0 {
		return nil, utils.DataNotFoundError(keyword)
	} else {
		fmt.Printf("Product found: [%v]\n", fnbs)
	}
	return fnbs, nil
}
func NewFindMenuUseCase(fnbRepo repository.FnBRepository) FindMenuUseCase {
	return FindMenuUseCase{
		fnbRepo: fnbRepo,
	}
}
