package product

import "fmt"

type SaleStatus string

const (
	SaleStatusOnSale  SaleStatus = "ON_SALE"
	SaleStatusSoldOut SaleStatus = "SOLD_OUT"
	SaleStatusStopped SaleStatus = "STOPPED"
)

func (s SaleStatus) IsValid() bool {
	switch s {
	case SaleStatusOnSale, SaleStatusSoldOut, SaleStatusStopped:
		return true
	}
	return false
}

func ParseSaleStatus(v string) (SaleStatus, error) {
	s := SaleStatus(v)
	if !s.IsValid() {
		return "", fmt.Errorf("invalid sale status: %s", v)
	}
	return s, nil
}
