package model

import (
	"time"
)

type Store struct {
	ID          uint      `gorm:"column:id;primaryKey;autoincrement" json:"id"`
	FullAddress string    `gorm:"type:varchar(255)" json:"full_address"`
	Latitude    string    `gorm:"type:varchar(255)" json:"latitude"`
	Longitude   string    `gorm:"type:varchar(255)" json:"logitude"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `gorm:"update_time_stamp" json:"updated_at"`
}

func (Store) TableName() string {
	return "store"
}
