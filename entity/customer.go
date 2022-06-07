package entity

import "fmt"

type Customer struct {
	CustomerId    string
	MobilePhoneNo string
	Name          string
}

func (c Customer) String() string {
	return fmt.Sprintf("Id: %s, Phone: %s, Name: %s", c.CustomerId, c.MobilePhoneNo, c.Name)
}
