package Handlers

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/taingk/goxit/api/Helpers"
	"github.com/taingk/goxit/api/Models"
)

func ListVote(c *gin.Context) {
	var vote []Models.Vote
	err := Models.GetAllVote(&vote)
	if err != nil {
		Helpers.RespondJSON(c, 404, vote)
	} else {
		Helpers.RespondJSON(c, 200, vote)
	}
}

func AddNewVote(c *gin.Context) {
	if Models.Authorize(c) == true {
		var vote Models.Vote
		vote.UUID = uuid.NewV4().String()
		c.BindJSON(&vote)
		err := Models.AddNewVote(&vote)
		if err != nil {
			Helpers.RespondJSON(c, 404, vote)
		} else {
			Helpers.RespondJSON(c, 200, vote)
		}
	} else {
		Helpers.RespondJSON(c, 404, "Access not granted")
	}
}

func GetOneVote(c *gin.Context) {
	uuid := c.Params.ByName("uuid")
	var vote Models.Vote
	err := Models.GetOneVote(&vote, uuid)
	if err != nil {
		Helpers.RespondJSON(c, 404, vote)
	} else {
		Helpers.RespondJSON(c, 200, vote)
	}
}

func PutOneVote(c *gin.Context) {
	if Models.Authorize(c) == true {
		var vote Models.Vote
		uuid := c.Params.ByName("uuid")
		err := Models.GetOneVote(&vote, uuid)
		if err != nil {
			Helpers.RespondJSON(c, 404, vote)
		}
		c.BindJSON(&vote)
		err = Models.PutOneVote(&vote, uuid)
		if err != nil {
			Helpers.RespondJSON(c, 404, vote)
		} else {
			Helpers.RespondJSON(c, 200, vote)
		}
	} else if Models.AuthorizeSimpleUser(c) == true {
		userUUID := c.Request.Header.Get("Authorization")
		voteUUID := c.Params.ByName("uuid")
		var vote Models.Vote
		err := Models.GetOneVote(&vote, voteUUID)
		if err != nil {
			Helpers.RespondJSON(c, 404, vote)
		}
		UUIDVotes := vote.UUIDVotes
		UUIDVotes = append(UUIDVotes, userUUID)
		vote.UUIDVotes = UUIDVotes
		err = Models.PutOneVote(&vote, voteUUID)
		if err != nil {
			Helpers.RespondJSON(c, 404, vote)
		} else {
			Helpers.RespondJSON(c, 200, vote)
		}
	} else {
		Helpers.RespondJSON(c, 404, "Access not granted")
	}
}

func DeleteVote(c *gin.Context) {
	if Models.Authorize(c) == true {
		var vote Models.Vote
		uuid := c.Params.ByName("uuid")
		err := Models.DeleteVote(&vote, uuid)
		if err != nil {
			Helpers.RespondJSON(c, 404, vote)
		} else {
			Helpers.RespondJSON(c, 200, vote)
		}
	} else {
		Helpers.RespondJSON(c, 404, "Access not granted")
	}
}
