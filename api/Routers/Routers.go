package Routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/taingk/goxit/api/Handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,AccessLevel"},
		MaxAge:       4 * time.Hour,
	}))
	r.POST("login", Handlers.Login)
	r.GET("user", Handlers.ListUser)
	r.POST("user", Handlers.AddNewUser)
	r.GET("user/:uuid", Handlers.GetOneUser)
	r.PUT("user/:uuid", Handlers.PutOneUser)
	r.DELETE("user/:uuid", Handlers.DeleteUser)
	r.GET("vote", Handlers.ListVote)
	r.POST("vote", Handlers.AddNewVote)
	r.GET("vote/:uuid", Handlers.GetOneVote)
	r.PUT("vote/:uuid", Handlers.PutOneVote)
	r.DELETE("vote/:uuid", Handlers.DeleteVote)

	return r
}
