package Routers

import (
	"github.com/gin-gonic/gin"
	"github.com/taingk/goxit/api/Handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("user", Handlers.ListUser)
	r.POST("user", Handlers.AddNewUser)
	r.GET("user/:id", Handlers.GetOneUser)
	r.PUT("user/:id", Handlers.PutOneUser)
	r.DELETE("user/:id", Handlers.DeleteUser)
	r.GET("vote", Handlers.ListVote)
	r.POST("vote", Handlers.AddNewVote)
	r.GET("vote/:id", Handlers.GetOneVote)
	r.PUT("vote/:id", Handlers.PutOneVote)
	r.DELETE("vote/:id", Handlers.DeleteVote)

	return r
}
