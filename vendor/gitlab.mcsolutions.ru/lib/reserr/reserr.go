package reserr

import "errors"

//logins
var (
	LOGIN_NOT_DEFINED       = errors.New("login not defined")
	ADMIN_LOGIN_NOT_DEFINED = errors.New("admin login not defined")
	LOGIN_NOT_FOUND         = errors.New("login not found")
	LOGIN_NOT_ACTIVE        = errors.New("login not active")
	PASSWORD_NOT_MATCHED    = errors.New("password not matched")
	NO_ONE_ROLE             = errors.New("no one role")
)

//Tokens
var (
	BAD_AUTHORIZATION_TOKEN = errors.New("bad authorization token")
)

//database
var (
	DATABASE_CONNECTION_ERROR = errors.New("database connection error")
	DATABASE_ERROR            = errors.New("database error")
)
