package database

import (
	"database/sql"
)

type UserId = int

type UserData struct {
	Email    string
	Password string
	Salt     string
}

type User struct {
	Id   UserId
	Data UserData
}

func (s *service) CreateUser(data UserData) (sql.Result, error) {
	const query = `INSERT INTO user (email, password, salt) VALUES ($email, $password, $salt)`
	statement, _ := s.db.Prepare(query)
	return statement.Exec(data.Email, data.Password, data.Salt)
}
