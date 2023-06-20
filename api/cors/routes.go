package cors

import (
	"github.com/gin-gonic/gin"

	"github.com/lanngoen1996/golang-basic/service/healthz"
	"github.com/lanngoen1996/golang-basic/service/users"
)

func Cors(r *gin.Engine) {
	healthz.Router(r)
	users.Router(r)
}
