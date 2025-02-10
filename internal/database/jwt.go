package database

import "database/sql"

type Jwt struct {
	UserId UserId
	Token  string
}

func (s *service) CreateJwt(data Jwt) (sql.Result, error) {
	const query = "INSERT INTO jwt (jwt, user_id) VALUES($jwt, $user_id)"
	statement, _ := s.db.Prepare(query)
	return statement.Exec(data.Token, data.UserId)
}
