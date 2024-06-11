package controller

import (
	"acana/dao/mysql"
	"acana/logic"
	"acana/models"
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// step1: get and check params
	p := new(models.ParamsSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Signup with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// check params
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || (p.Password != p.RePassword) {
		zap.L().Error("Signup with invalid params")
		ok, _ := mysql.CheckUserIsExist(p.Username)
		if !ok {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
	}

	// step2: handle business
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// step3: return reponses
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// step1: get requset params and check params
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// step2: handle the busniess
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login.Login failed", zap.String("username:", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// step3: return reponse
	ResponseSuccess(c, gin.H{
		"user_id":   user.UserID,
		"user_name": user.Username,
		"token":     user.Token,
	})
}
