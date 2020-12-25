package databases
import(
	"log"
	m "com/tienxe/lib/server/user_service/models"
	sql "database/sql"
)
type MySqlAdapter struct{
	db     *sql.DB
}

func NewMysql(sqlPath string) (*MySqlAdapter){
	DB, err := sql.Open("mysql",sqlPath)
	log.Printf("error %v",err)
	_, err = DB.Exec("use microservice")
	log.Printf("error use database %v",err)

	_, err = DB.Exec("create table users(id int auto_increment primary key,cardnumber varchar(32),Name varchar(64),Age int)")
	log.Printf("error create tables %v",err)

	return &MySqlAdapter{db: DB}
}
func (a *MySqlAdapter)InsertUser(u *m.User) error{
	a.db.Exec("insert into users(name, age) values(?,?)",u.Name, u.Age)
	return nil
}
func (a *MySqlAdapter)UpdateUser(u *m.User, id string) error{
	a.db.Exec("update users set name=?,age=? where id=?", u.Name, u.Age, id)
	return nil
}
func (a *MySqlAdapter)GetUser() ([]m.User, error){
	rows, err := a.db.Query("select * from users")
	if err == nil{
		results :=  make([]m.User,0)
		for rows.Next(){
			item := m.User{}
			if err := rows.Scan(&item.Name,&item.Age, &item.CardNumber); err != nil{
				//fmt.Printf("Scan error %v",err)
				continue
			}
			results = append(results, item)
	
		}
		return results, nil
	}
	
	return nil, nil
}
func (a *MySqlAdapter)FindUser(id string) (m.User,error){
	rows := a.db.QueryRow("select * from users where cardnumber=?",id)
	result := m.User{}
	if err := rows.Scan(&result.CardNumber, &result.Name, &result.Age); err != nil{
		return result, err
	}
	return result, nil
}
func (a *MySqlAdapter)DeleteUser(id string) error{
	a.db.Exec("delete from users where id=?",id)
	return nil
}