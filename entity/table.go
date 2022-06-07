package entity

import "fmt"

type Table struct {
	TableNo     string
	IsAvailable bool
}

func (t Table) String() string {
	return fmt.Sprintf("TableNo: %s, Available: %v", t.TableNo, t.IsAvailable)
}
