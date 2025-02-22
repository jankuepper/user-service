package database

import (
	"database/sql"
)

type SerieId = int
type SerieData struct {
	Name          string
	ThumbnailPath string
}

type Serie struct {
	Id   SerieId
	Data SerieData
}

func (s *service) CreateSerie(data SerieData) (sql.Result, error) {
	const query = `INSERT INTO serie (name, thumbnailpath) VALUES ($name, $thumbnailpath)`
	statement, _ := s.db.Prepare(query)
	return statement.Exec(data.Name, data.ThumbnailPath)
}

func (s *service) GetAllSeries() ([]Serie, error) {
	rows, err := s.db.Query("SELECT * FROM serie")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	series := []Serie{}
	for rows.Next() {
		var serie Serie
		if err = rows.Scan(&serie.Id, &serie.Data.Name, &serie.Data.ThumbnailPath); err != nil {
			return nil, err
		}
		series = append(series, serie)
	}
	return series, err
}
