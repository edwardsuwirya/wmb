package entity

import "fmt"

type Table struct {
	TableNo string
}

func (t Table) String() string {
	return fmt.Sprintf("Id: %s", t.TableNo)
}
