package cors

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/lanngoen1996/golang-basic/service/healthz"
	"github.com/lanngoen1996/golang-basic/service/users"
)

func NewApp(r *gin.Engine) {
	db := Database()
	router(r, db)

	/* migration */
	Migration(db)
}

func router(r *gin.Engine, db *gorm.DB) {
	group := r.Group("api/v1")

	healthz.Router(group)
	users.Router(group, db)
}
