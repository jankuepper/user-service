package database

import (
	"database/sql"
	"fmt"
)

// Serie hat Seasons, Seasons haben Episoden

type SerieData struct {
	Name          string
	ThumbnailPath string
}

type Serie struct {
	Id   UserId
	Data SerieData
}

func (s *service) CreateSerie(data SerieData) (sql.Result, error) {
	const query = `INSERT INTO serie (name, thumbnailpath) VALUES ($name, $thumbnailpath)`
	statement, _ := s.db.Prepare(query)
	return statement.Exec(data.Name, data.ThumbnailPath)
}

func (s *service) GetAllSeries() ([]Serie, error) {
	query := `SELECT * FROM serie`
	rows, err := s.db.Query(query)
	if err != nil {
		var serie Serie
		return []Serie{serie}, err
	}
	defer rows.Close()
	series := []Serie{}
	for rows.Next() {
		var serie Serie
		if err = rows.Scan(&serie.Id, &serie.Data.Name, &serie.Data.ThumbnailPath); err != nil {
			return series, err
		}
		series = append(series, serie)
	}
	fmt.Println("Serie ", series)
	return series, err
}
