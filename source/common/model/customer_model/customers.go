package customermodel

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          uint           `gorm:"primaryKey" json:"-"`
	FirstName   string         `gorm:"not null;size:100" json:"first_name"`
	LastName    string         `gorm:"not null;size:100" json:"last_name"`
	Email       string         `gorm:"unique;not null;size:255" json:"email"`
	Phone       string         `gorm:"size:20" json:"phone,omitempty"`
	Address     string         `gorm:"type:text" json:"address,omitempty"`
	City        string         `gorm:"size:100" json:"city,omitempty"`
	Country     string         `gorm:"size:100;default:'Indonesia'" json:"country"`
	DateOfBirth *time.Time     `gorm:"type:date" json:"date_of_birth,omitempty"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName returns the table name for Customer model
func (Customer) TableName() string {
	return "customers"
}

// FullName returns concatenated first and last name
func (c Customer) FullName() string {
	return c.FirstName + " " + c.LastName
}
