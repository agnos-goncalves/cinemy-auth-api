package model

type User struct {
	Id     int
	Email  string
	Active bool
}

type UserFull struct {
	Id       int
	Email    string
	Active   bool
	Password string
}