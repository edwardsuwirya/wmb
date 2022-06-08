package utils

import "fmt"

type TableUnavailableError struct {
	TableNo string
}

func (e TableUnavailableError) Error() string {
	return fmt.Sprintf("Table %s is not available\n", e.TableNo)
}

type ProductNotFoundError struct {
	ProductInfo string
}

func (e ProductNotFoundError) Error() string {
	return fmt.Sprintf("Product [%s] not found\n", e.ProductInfo)
}

type BillNotFoundError struct {
	BillNo string
}

func (e BillNotFoundError) Error() string {
	return fmt.Sprintf("Bill [%s] not found\n", e.BillNo)
}
