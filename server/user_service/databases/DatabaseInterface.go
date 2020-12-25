package databases
import(
	m "com/tienxe/lib/server/user_service/models"
)

type Adapter interface{
	InsertUser(u *m.User) error
	UpdateUser(u *m.User, id string) error
	GetUser() ([]m.User, error)
	FindUser(id string) (m.User, error)
	DeleteUser(id string) error
}