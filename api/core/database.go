package core

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		GetEnv().DB_USERNAME,
		GetEnv().DB_PASSWORD,
		GetEnv().DB_HOST,
		GetEnv().DB_PORT,
		GetEnv().DB_NAME,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database !")
	}

	DB = conn
}
