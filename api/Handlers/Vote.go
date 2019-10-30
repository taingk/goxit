package Handlers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/taingk/goxit/api/ApiHelpers"
	"github.com/taingk/goxit/api/Models"
)

func ListVote(c *gin.Context) {
	var vote []Models.Vote
	err := Models.GetAllVote(&vote)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	} else {
		ApiHelpers.RespondJSON(c, 200, vote)
	}
}

func AddNewVote(c *gin.Context) {
	var vote Models.Vote
	vote.UUID = uuid.NewV4().String()
	c.BindJSON(&vote)
	err := Models.AddNewVote(&vote)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	} else {
		ApiHelpers.RespondJSON(c, 200, vote)
	}
}

func GetOneVote(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	var vote Models.Vote
	err := Models.GetOneVote(&vote, uuid)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	} else {
		ApiHelpers.RespondJSON(c, 200, vote)
	}
}

func PutOneVote(c *gin.Context) {
	var vote Models.Vote
	uuid := c.Params.ByName("uuid")
	err := Models.GetOneVote(&vote, uuid)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	}
	c.BindJSON(&vote)
	err = Models.PutOneVote(&vote, uuid)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	} else {
		ApiHelpers.RespondJSON(c, 200, vote)
	}
}

func DeleteVote(c *gin.Context) {
	var vote Models.Vote
	uuid := c.Params.ByName("uuid")
	err := Models.DeleteVote(&vote, uuid)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	} else {
		ApiHelpers.RespondJSON(c, 200, vote)
	}
}
