package authorization
import("database/sql"
_ "github.com/go-sql-driver/mysql")
type AuthenHandler struct{
	//db object
	db *sql.DB

}
func NewAuthenHandler(mysql string) AuthService{
	//create db object
	Db, err := sql.Open("mysql",mysql)
	if(err != nil){
		panic(err.Error())
	}
	return &AuthenHandler{db: Db}
}

//Add
func (a * AuthenHandler) RegisterUser(user* AuthUser) error{
	return nil
}
//Login
func (a * AuthenHandler) LoginUser(user* AuthUser) error{
	return nil
}
//Update Password
func (a * AuthenHandler) UpdateUser(user* AuthUser)error{
	return nil

}
//delete
func (a * AuthenHandler) DeleteUser(id string)error{
	return nil
}