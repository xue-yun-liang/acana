package controller

import (
	"acana/logic"
	"acana/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// type VoteData struct {
// 	PostID    int64 `json:"post_id, string"`
// 	Direction int   `json:"direction, string"`
// }

func PostVoteHandler(c *gin.Context) {
	// step1: get and check params
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		return
	}

	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// the business logic for specific votes
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
