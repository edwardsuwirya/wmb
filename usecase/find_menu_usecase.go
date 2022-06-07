package usecase

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"fmt"
)

type FindMenuUseCase struct {
	fnbRepo repository.FnBRepository
}

func (c *FindMenuUseCase) FindMenuById(id string) entity.FnB {
	fnb := c.fnbRepo.FindById(id)
	if fnb.FnBId == "" {
		fmt.Printf("Product with ID %s not found\n", fnb.FnBId)
	} else {
		fmt.Printf("Product found : [%v]\n", fnb)
	}
	return fnb
}
func (c *FindMenuUseCase) FindMenuByName(keyword string) []entity.FnB {
	fnbs := c.fnbRepo.FindByName(keyword)
	if len(fnbs) == 0 {
		fmt.Printf("Product with keyword %s not found\n", keyword)
	} else {
		fmt.Printf("Product found: [%v]\n", fnbs)
	}
	return fnbs
}
func NewFindMenuUseCase(fnbRepo repository.FnBRepository) FindMenuUseCase {
	return FindMenuUseCase{
		fnbRepo: fnbRepo,
	}
}
