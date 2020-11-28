package authorization

type AuthService interface{
	//Add
	RegisterUser(user* AuthUser) error
	//Login
	LoginUser(user* AuthUser) error
	//Update Password
	UpdateUser(user* AuthUser)error
	//delete
	DeleteUser(id string)error
}