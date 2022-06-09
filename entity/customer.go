package entity

import (
	"fmt"
	"time"
)

type CustomerPrivileges struct {
	DiscountPct float64
}

type Customer struct {
	CustomerId    string
	MobilePhoneNo string
	Name          string
	ActiveMember  bool
	JoinDate      time.Time
	CustomerPrivileges
}

func (c Customer) String() string {
	return fmt.Sprintf("Id: %s, Phone: %s, Name: %s, Active Member: %v, Join Date: %v, Privileges: %v", c.CustomerId, c.MobilePhoneNo, c.Name, c.ActiveMember, c.JoinDate, c.CustomerPrivileges)
}
