package logic

import (
	"acana/dao/mysql"
	"acana/models"
	"errors"

	"acana/pkg/jwt"
	"acana/pkg/snowflake"
)

func SignUp(p *models.ParamsSignUp) (err error) {
	// step1: check user name already exist?
	var is_exist bool
	is_exist, err = mysql.CheckUserIsExist(p.Username)
	if err != nil {
		// mysql query error
		return err
	}
	if is_exist {
		return errors.New("user has exist, please change name")
	}

	mysql.QueryUserByUsername()
	// step2: generate user id
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// step3: store user into database
	mysql.InsertUser(user)
	return
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// pass user's pointer
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// Generate jwt token
	return jwt.GenToken(user.UserID, user.Username)
}
