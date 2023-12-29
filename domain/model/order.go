package model

import "time"

type Order struct {
	ID             uint      `gorm:"column:id;primaryKey;autoincrement" json:"id"`
	UserID         uint      `gorm:"not null"`
	User           User      `gorm:"foreignKey:UserID"`
	ProductID      uint      `gorm:"not null"`
	Product        Product   `gorm:"foreignKey:ProductID"`
	Quantity       int       `gorm:"not null"`
	PurchaseScheme string    `gorm:"not null"`
	PaymentStatus  string    `gorm:"not null"`
	DeliveryStatus string    `gorm:"not null"`
	DeliveryDate   time.Time `gorm:"default:null"`
	SupplierID     uint      `gorm:"not null"`
	Supplier       Supplier  `gorm:"foreignKey:SupplierID"`
}

func (Order) TableName() string {
	return "order"
}
