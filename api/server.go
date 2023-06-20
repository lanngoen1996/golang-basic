package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lanngoen1996/golang-basic/cors"
)

func main() {
	r := gin.Default()

	cors.Cors(r)

	/* listen and serve on 0.0.0.0:8080 */

	r.Run(cors.GetEnv().HOST + ":" + cors.GetEnv().PORT)
}
