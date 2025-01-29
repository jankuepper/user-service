package database

import "log"

type UserId = int

type UserData struct {
	email    string
	password string
	salt     string
}

type User struct {
	id   UserId
	data UserData
}

func (s *service) createUser(data UserData) {
	const query = `INSERT INTO user (email, password, type, salt) VALUES ($email, $password, $type, $salt)`
	statement, _ := s.db.Prepare(query)
	_, err := statement.Exec(data.email, data.password, data.salt)
	if err != nil {
		log.Printf("Error in creating user with email:%s\n", data.email)
		return
	}
	log.Println("Successfully updated the book in database!")
}
