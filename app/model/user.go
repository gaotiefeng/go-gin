package model

import db "go-gin/database"

type User struct {
	Id			int `json:"id"`
	Mobile		string	`json:"mobile"`
	Password	string	`json:"password"`
}

func (u *User) AddUser()	(id int64, err error)  {

	rs, err := db.SqlDB.Exec("INSERT INTO user(mobile, password) values (?,?)", u.Mobile, u.Password)

	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
