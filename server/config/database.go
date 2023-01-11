package config

import (
	"github.com/jossep11/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:@tcp(127.0.0.1:3306)/dog?parseTime=true"

func Connect() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Data{})
	Database.AutoMigrate(&entities.Users{})

	return nil
}
