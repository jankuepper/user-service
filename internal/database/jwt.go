package database

import "log"

type Jwt struct {
	userId UserId
	token  string
}

func (s *service) createJwt(data Jwt) {
	const query = "" // TODO `INSERT INTO jwt (email, password, type, salt) VALUES ($email, $password, $type, $salt)`
	statement, _ := s.db.Prepare(query)
	_, err := statement.Exec(data.token)
	if err != nil {
		log.Printf("Error in creating jwt %s for user with id:%d\n", data.token, data.userId)
		return
	}
	log.Println("Successfully updated the book in database!")
}
