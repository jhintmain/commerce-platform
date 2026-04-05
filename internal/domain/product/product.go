package product

import "time"

type Product struct {
	ID          int64
	BrandID     int64
	Name        string
	Description string
	BasePrice   int
	SaleStatus  SaleStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// 판매중
func (p *Product) IsOnSale() bool {
	return p.SaleStatus == SaleStatusOnSale
}

// 품절
func (p *Product) IsSoldOut() bool {
	return p.SaleStatus == SaleStatusSoldOut
}

// 판매중단
func (p *Product) IsStopped() bool {
	return p.SaleStatus == SaleStatusStopped
}

// 삭제됨
func (p *Product) IsDeleted() bool {
	return p.DeletedAt != nil
}
