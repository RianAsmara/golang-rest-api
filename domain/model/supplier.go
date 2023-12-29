package model

import "time"

type Supplier struct {
	ID        uint   `gorm:"column:id;primaryKey;autoincrement" json:"id"`
	Name      string `gorm:"type:varchar(50);not null"`
	Email     string `gorm:"type:varchar(100);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	StoreID   uint
	Store     Store `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"update_time_stamp"`
}

func (Supplier) TableName() string {
	return "supplier"
}
