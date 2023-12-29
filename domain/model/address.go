package model

type Address struct {
	ID           uint     `gorm:"column:id;primaryKey;autoincrement" json:"id"`
	AddressLine1 string   `gorm:"type:varchar(100);not null"`
	AddressLine2 string   `gorm:"type:varchar(100)"`
	City         string   `gorm:"type:varchar(50);not null"`
	State        string   `gorm:"type:varchar(50);not null"`
	ZipCode      string   `gorm:"type:varchar(10);not null"`
	UserID       uint     `gorm:"null"`
	User         User     `gorm:"foreignKey:UserID"`
	SupplierID   uint     `gorm:"null"`
	Supplier     Supplier `gorm:"foreignKey:SupplierID"`
}

func (Address) TableName() string {
	return "address"
}
