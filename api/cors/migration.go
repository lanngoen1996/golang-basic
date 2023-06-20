package cors

import (
	"gorm.io/gorm"

	"github.com/lanngoen1996/golang-basic/service/users"
)

func Migration(DB *gorm.DB) {
	DB.AutoMigrate(&users.User{})
}
