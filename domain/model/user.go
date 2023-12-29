package model

import "time"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	ID          uint      `gorm:"column:id;primaryKey;autoincrement" json:"id"`
	Name        string    `gorm:"type:varchar(50);not null" json:"name"`
	Email       string    `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"password"`
	FullAddress string    `gorm:"type:varchar(255)" json:"full_address"`
	Latitude    string    `gorm:"type:varchar(255)" json:"latitude"`
	Longitude   string    `gorm:"type:varchar(255)" json:"longitude"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `gorm:"update_time_stamp" json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
