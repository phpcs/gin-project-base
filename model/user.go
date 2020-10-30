package model

type Users struct{
	Id  int `json:"id"`
	Username string `json:"user_name"`
	Password string `json:"-"`
	Email string	`json:"email"`
}

func (u *Users) GetList(id int) []Users{
	var arr []Users
	db.Raw("SELECT * FROM users where id=?", id).Scan(&arr)
	return arr
}