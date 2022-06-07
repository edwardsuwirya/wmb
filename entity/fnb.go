package entity

import "fmt"

type FnB struct {
	FnBId    string
	MenuName string
	Price    float64
}

func (fb FnB) String() string {
	return fmt.Sprintf("Id: %s, Menu: %v, Price: %f", fb.FnBId, fb.MenuName, fb.Price)
}
