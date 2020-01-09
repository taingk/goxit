package Handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/taingk/goxit/api/Helpers"
	"github.com/taingk/goxit/api/Models"
)

type Jwt struct {
	JWT         string
	AccessLevel int
}

func Login(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.GetUserByEmailPassword(&user)
	if err != nil {
		Helpers.RespondJSON(c, 401, user)
	} else {
		var jwt Jwt

		jwt.JWT = user.UUID
		jwt.AccessLevel = user.AccessLevel

		Helpers.RespondJSON(c, 200, jwt)
	}
}
