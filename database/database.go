package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dialectors = map[string]gorm.Dialector{
	"macapa": mysql.Open("root:admin@tcp(mysql:3306)/admin?charset=utf8mb4&parseTime=True&loc=Local"),
	"varejao": postgres.Open("host=postgresql user=admin password=admin dbname=admin port=5432"),
}

func GetDatabaseConnection(client string) (connection *gorm.DB, err error) {
	connection, err = gorm.Open(dialectors[client],&gorm.Config{})
	return
}

