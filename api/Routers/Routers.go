package Routers

import (
	"github.com/taingk/goxit/api/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("user", Controllers.ListUser)
	r.POST("user", Controllers.AddNewUser)
	r.GET("user/:id", Controllers.GetOneUser)
	r.PUT("user/:id", Controllers.PutOneUser)
	r.DELETE("user/:id", Controllers.DeleteUser)

	return r
}
