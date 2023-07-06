package model

type UserCreateInput struct {
	UserName string
	Password string
	Nickname string
}

type UserSignInInput struct {
	UserName string
	Password string
}
