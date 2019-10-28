package Handlers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/taingk/goxit/api/ApiHelpers"
	"github.com/taingk/goxit/api/Models"
	"github.com/gin-gonic/gin"
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
	id := c.Params.ByName("id")
	var vote Models.Vote
	err := Models.GetOneVote(&vote, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	} else {
		ApiHelpers.RespondJSON(c, 200, vote)
	}
}

func PutOneVote(c *gin.Context) {
	var vote Models.Vote
	id := c.Params.ByName("id")
	err := Models.GetOneVote(&vote, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	}
	c.BindJSON(&vote)
	err = Models.PutOneVote(&vote, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	} else {
		ApiHelpers.RespondJSON(c, 200, vote)
	}
}

func DeleteVote(c *gin.Context) {
	var vote Models.Vote
	id := c.Params.ByName("id")
	err := Models.DeleteVote(&vote, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, vote)
	} else {
		ApiHelpers.RespondJSON(c, 200, vote)
	}
}
