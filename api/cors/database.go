package cors

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() (conn *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		GetEnv().DB_USERNAME,
		GetEnv().DB_PASSWORD,
		GetEnv().DB_HOST,
		GetEnv().DB_PORT,
		GetEnv().DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database !")
	}

	return db
}
