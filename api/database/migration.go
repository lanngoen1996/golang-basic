package database

import (
	"github.com/lanngoen1996/golang-basic/core"
	model "github.com/lanngoen1996/golang-basic/models"
)

func Migration() {
	core.DB.AutoMigrate(&model.User{})
}
