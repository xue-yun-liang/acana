package mysql

import (
	"acana/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

/* warpper every opration of database, util the `logic` call */

const secret = "knsgwiu82jqy[;[_]]"

func QueryUserByUsername() {

}

var (
	ErrorUserExist       = errors.New("user have been exist")
	ErrorUserNotExist    = errors.New("user is not exist")
	ErrorInvalidPassword = errors.New("incorrect password")
)

func CheckUserIsExist(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var cnt int
	if err := db.Get(&cnt, sqlStr, username); err != nil {
		return false, err
	}
	return cnt > 0, nil
}

// InsertUser: insert a new user into mysql database
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	sqlStr := `insert into user(user_id, username, password) value(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(mPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(mPassword)))
}

func Login(user *models.User) (err error) {
	originPassword := user.Password
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return errors.New("no exist user")
	}
	if err != nil {
		return err
	}
	// check password
	password := encryptPassword(originPassword)
	if password != user.Password {
		return errors.New("Incorrect password")
	}
	return
}
