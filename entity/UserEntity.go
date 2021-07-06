package entity

import (
	DB "radiant-protal-api/database"
	"radiant-protal-api/models"
)

type LoginResult struct {
	ID              int    `json:"id"`
	UserCode        string `json:"user_code"`
	UserName        string `json:"user_name"`
	DesignationID   int    `json:"designation_id"`
	DesignationName string `json:"designation_name"`
	DateOfJoining   string `json:"date_of_joining"`
	DepartmentID    int    `json:"department_id"`
	DepartmentName  string `json:"department_name"`
	CompanyID       int    `json:"company_id"`
	CompanyName     string `json:"company_name"`
	JobTypeID       int    `json:"job_type_id"`
	CellPhone       string `json:"cell_phone"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Status          int    `json:"status"`
	Token           string `json:"token"`
}

func Login(userCode string, password string, userResult *[]LoginResult) (err error) {

	if err = DB.GetDB().Model(&models.UserListModel{}).
		Select("user_list.*, designation.designation_name, department.department_name, company.company_name").
		Joins("INNER JOIN designation ON user_list.designation_id = designation.id ").
		Joins("INNER JOIN department ON user_list.department_id = department.id ").
		Joins("INNER JOIN company ON user_list.company_id = company.id").
		Where("user_list.user_code = ? AND user_list.password = ?", userCode, password).
		Scan(userResult).Error; err != nil {
		return err
	}
	return nil

}

func CheckUser(userResult *[]LoginResult, UserCode string) (err error) {

	if err = DB.GetDB().Model(&models.UserListModel{}).
		Select("user_list.*, designation.designation_name, department.department_name, company.company_name").
		Joins("INNER JOIN designation ON user_list.designation_id = designation.id ").
		Joins("INNER JOIN department ON user_list.department_id = department.id ").
		Joins("INNER JOIN company ON user_list.company_id = company.id").
		Where("user_list.user_code = ?", UserCode).
		Scan(userResult).Error; err != nil {
		return err
	}
	return nil

}

func UserDetails(userResult *[]LoginResult, id int) (err error) {

	if err = DB.GetDB().Model(&models.UserListModel{}).
		Select("user_list.*, designation.designation_name, department.department_name, company.company_name").
		Joins("INNER JOIN designation ON user_list.designation_id = designation.id ").
		Joins("INNER JOIN department ON user_list.department_id = department.id ").
		Joins("INNER JOIN company ON user_list.company_id = company.id").
		Where("user_list.id = ?", id).
		Scan(userResult).Error; err != nil {
		return err
	}
	return nil

}
