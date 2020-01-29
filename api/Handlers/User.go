package Handlers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/taingk/goxit/api/Helpers"
	"github.com/taingk/goxit/api/Models"
)

func ListUser(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUser(&user)
	if err != nil {
		Helpers.RespondJSON(c, 404, user)
	} else {
		Helpers.RespondJSON(c, 200, user)
	}
}

func AddNewUser(c *gin.Context) {
	var user Models.User
	user.UUID = uuid.NewV4().String()
	c.BindJSON(&user)
	err := Models.AddNewUser(&user)
	if err != nil {
		Helpers.RespondJSON(c, 401, err)
	} else {
		Helpers.RespondJSON(c, 200, user)
	}
}

func GetOneUser(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	var user Models.User
	err := Models.GetOneUser(&user, uuid)
	if err != nil {
		Helpers.RespondJSON(c, 404, user)
	} else {
		Helpers.RespondJSON(c, 200, user)
	}
}

func PutOneUser(c *gin.Context) {
	if Models.SimpleAuth(c) == true {
		var user Models.User
		uuid := c.Params.ByName("uuid")
		err := Models.GetOneUser(&user, uuid)
		if err != nil {
			Helpers.RespondJSON(c, 404, user)
		}
		c.BindJSON(&user)
		err = Models.PutOneUser(&user, uuid)
		if err != nil {
			Helpers.RespondJSON(c, 401, user)
		} else {
			Helpers.RespondJSON(c, 200, user)
		}
	} else {
		Helpers.RespondJSON(c, 401, "Access not granted")
	}
}

func DeleteUser(c *gin.Context) {
	if Models.SimpleAuth(c) == true {
		var user Models.User
		uuid := c.Params.ByName("uuid")
		err := Models.GetOneUser(&user, uuid)
		if err != nil {
			Helpers.RespondJSON(c, 404, user)
		}
		c.BindJSON(&user)
		errdel := Models.DeleteUser(&user, uuid)
		if errdel != nil {
			Helpers.RespondJSON(c, 404, user)
		} else {
			Helpers.RespondJSON(c, 200, user)
		}
	} else {
		Helpers.RespondJSON(c, 401, "Access not granted")
	}
}
