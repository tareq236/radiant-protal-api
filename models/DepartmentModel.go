package models

import "time"

type DepartmentModel struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	DepartmentName string    `json:"department_name"`
	Status         int       `gorm:"default:1"`
	CreatedAt      time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *DepartmentModel) TableName() string {
	return "department"
}
