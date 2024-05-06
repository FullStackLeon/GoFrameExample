package consts

import "errors"

var (
	ErrCodeSuccess         = 0
	ErrCodeAccountNotExist = 100
	ErrCodeUserRegistered  = 101

	ErrCodeQueryUserFailed = 102
	ErrCodeRegisterFailed  = 103
)

var (
	ErrMsgSuccess         = errors.New("success")
	ErrMsgAccountNotExist = errors.New("account not exist")
	ErrMsgUserRegistered  = errors.New("user has been registered")

	ErrMsgQueryUserFailed = errors.New("query user failed")
	ErrMsgRegisterFailed  = errors.New("register failed")

	ErrMsgGenerateTokenFailed = errors.New("generate token failed")
	ErrMsgSetSessionFailed    = errors.New("set session failed")
)
