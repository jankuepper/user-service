package database

import (
	"log"
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

func (s *service) CreateUser(data UserData) {
	const query = `INSERT INTO user (email, password, type, salt) VALUES ($email, $password, $type, $salt)`
	statement, _ := s.db.Prepare(query)
	_, err := statement.Exec(data.Email, data.Password, data.Salt)
	if err != nil {
		log.Printf("Error in creating user with email:%s\n", data.Email)
		return
	}
	log.Println("Successfully updated the book in database!")
}
