package usecase

import (
	"enigmacamp.com/wmbpos/entity"
	"enigmacamp.com/wmbpos/repository"
	"enigmacamp.com/wmbpos/utils"
	"fmt"
	"time"
)

type MembershipUseCase struct {
	customerRepo repository.CustomerRepository
}

func (m *MembershipUseCase) RegisterNewCustomer(customer entity.Customer) (string, error) {
	if customer.MobilePhoneNo == "" {
		return "", utils.RequiredError("Customer Phone")
	}
	cust := m.customerRepo.FindByPhone(customer.MobilePhoneNo)
	if cust.CustomerId != "" {
		return m.RegisterExistingCustomer(customer)
	} else {
		customer.ActiveMember = true
		customer.JoinDate = time.Now()
		customer.DiscountPct = 0.1
		m.customerRepo.Create(&customer)
		fmt.Printf("New member activation: [%v] \n", customer.CustomerId)
		return customer.CustomerId, nil
	}
}
func (m *MembershipUseCase) RegisterExistingCustomer(customer entity.Customer) (string, error) {
	if customer.MobilePhoneNo == "" {
		return "", utils.RequiredError("Customer Phone")
	}
	cust := m.customerRepo.FindByPhone(customer.MobilePhoneNo)
	if cust.CustomerId == "" {
		return "", utils.DataNotFoundError(customer.MobilePhoneNo)
	}
	if cust.ActiveMember {
		return "", utils.AlreadyExistError(cust.CustomerId)
	}
	cust.ActiveMember = true
	cust.JoinDate = time.Now()
	cust.DiscountPct = 0.1
	m.customerRepo.Update(&cust)
	fmt.Printf("New member activation: [%v] \n", cust)
	return customer.CustomerId, nil
}

func NewMembershipUseCase(customerRepo repository.CustomerRepository) MembershipUseCase {
	return MembershipUseCase{
		customerRepo: customerRepo,
	}
}
