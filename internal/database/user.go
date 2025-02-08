package database

import (
	"database/sql"
)

type UserId = int

type UserData struct {
	Email    string
	Password string
}

type User struct {
	Id   UserId
	Data UserData
}

func (s *service) CreateUser(data UserData) (sql.Result, error) {
	const query = `INSERT INTO user (email, password) VALUES ($email, $password)`
	statement, _ := s.db.Prepare(query)
	return statement.Exec(data.Email, data.Password)
}

func (s *service) GetUserByEmail(email string) (User, error) {
	var user User
	const query = `SELECT id, email, password FROM user WHERE email = $email`
	rows, err := s.db.Query(query, email)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.Id, &user.Data.Email, &user.Data.Password)
	}
	return user, err
}
