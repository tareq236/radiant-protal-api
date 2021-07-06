package models

import "time"

type UserListModel struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserCode      string    `json:"user_code"`
	UserName      string    `json:"user_name"`
	DesignationID int       `gorm:"default:NULL" json:"designation_id"`
	DateOfJoining string    `gorm:"default:NULL" json:"date_of_joining"`
	DepartmentID  int       `gorm:"default:NULL" json:"department_id"`
	CompanyID     int       `gorm:"default:NULL" json:"company_id"`
	JobTypeID     int       `gorm:"default:NULL" json:"job_type_id"`
	CellPhone     string    `gorm:"default:NULL" json:"cell_phone"`
	Email         string    `gorm:"default:NULL" json:"email"`
	Password      string    `gorm:"default:NULL" json:"password"`
	Status        int       `gorm:"default:1"`
	CreatedAt     time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *UserListModel) TableName() string {
	return "user_list"
}
