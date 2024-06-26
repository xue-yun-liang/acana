package controller

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedLogin
	CodeInvalidToken
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "incorrect params",
	CodeUserExist:       "user have been exist",
	CodeUserNotExist:    "user is not exist",
	CodeInvalidPassword: "incorrect username or password",
	CodeServerBusy:      "server buys",

	CodeNeedLogin:    "need login",
	CodeInvalidToken: "invalid token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
