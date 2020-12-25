package daos
import (
	"log"
	d "com/tienxe/lib/server/user_service/databases"
	m "com/tienxe/lib/server/user_service/models"
)


type UserDao struct {
	a d.Adapter
}
func InitUserDao(config string)(dao * UserDao) {
	log.Printf("init user dao %v",config)
	instanse := d.NewMysql(config)
	return &UserDao{a :  instanse}
}
func(u* UserDao) AddUser(user* m.User) error{
	//call global database
	return u.a.InsertUser(user)
}
func(u* UserDao) FindUser(id string) (m.User,error){
	return u.a.FindUser(id)
}
func(u* UserDao) UpdateUser(user* m.User) error{
	return nil
}
func(u* UserDao) DeleteUser(user* m.User) error{
	return nil
}