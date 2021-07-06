package models

import "time"

type DesignationModel struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	DesignationName string    `json:"designation_name"`
	Status          int       `gorm:"default:1"`
	CreatedAt       time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *DesignationModel) TableName() string {
	return "designation"
}
