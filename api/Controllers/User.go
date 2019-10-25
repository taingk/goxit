package Controllers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/taingk/goxit/api/ApiHelpers"
	"github.com/taingk/goxit/api/Models"
	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUser(&user)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func AddNewUser(c *gin.Context) {
	var user Models.User
	user.UUID = uuid.NewV4().String()
	c.BindJSON(&user)
	err := Models.AddNewUser(&user)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func GetOneUser(c *gin.Context) {
	id := c.Params.ByName("UUID")
	var user Models.User
	err := Models.GetOneUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func PutOneUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("UUID")
	err := Models.GetOneUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	}
	c.BindJSON(&user)
	err = Models.PutOneUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("UUID")
	err := Models.DeleteUser(&user, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, user)
	} else {
		ApiHelpers.RespondJSON(c, 200, user)
	}
}
