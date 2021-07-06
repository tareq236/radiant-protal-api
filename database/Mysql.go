package database

import (
	"fmt"
	"os"
	"strconv"

	"radiant-protal-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var MYsqlDB *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Println(e)
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	if err != nil {
		fmt.Println("database port convert error:", err)
	}

	dbUri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	// fmt.Println(dbUri)

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Println("mysql database connection error", err)
	} else {
		fmt.Println("mysql database connected")
	}

	MYsqlDB = conn
	// defer db.Close()
	MYsqlDB.AutoMigrate(
		&models.TokenModel{},
		&models.CompanyModel{},
		&models.DepartmentModel{},
		&models.DesignationModel{},
		&models.JobTypeModel{},
		&models.UserListModel{})
	// db.Debug().AutoMigrate(&Account{}, &Contact{})

}

func GetDB() *gorm.DB {
	return MYsqlDB
}
