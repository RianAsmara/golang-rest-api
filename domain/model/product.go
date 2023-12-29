package model

import "github.com/shopspring/decimal"

type Product struct {
	ID             uint            `gorm:"column:id;primaryKey;autoincrement" json:"id"`
	Name           string          `gorm:"type:varchar(50);not null"`
	Description    string          `gorm:"type:varchar(255)"`
	Price          decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	PurchaseScheme string          `gorm:"not null"`
	SupplierID     uint            `gorm:"not null"`
	Supplier       Supplier        `gorm:"foreignKey:SupplierID"`
}

func (Product) TableName() string {
	return "product"
}
