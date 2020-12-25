package models

type User struct{
	Name string
	Age	 int
	CardNumber string
}
//model use for Authen
type UserAuthen struct{
	UserName string 
	PassWord string
}