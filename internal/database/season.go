package database

import "database/sql"

type SeasonId = int
type SeasonData struct {
	Name          string
	ThumbnailPath string
	SerieId       SerieId
}

type Season struct {
	Id   SeasonId
	Data SeasonData
}

func (s *service) CreateSeason(data SeasonData) (sql.Result, error) {
	const query = `INSERT INTO season (name, thumbnailpath, serieid) VALUES ($name, $thumbnailpath, $serieid)`
	statement, _ := s.db.Prepare(query)
	return statement.Exec(data.Name, data.ThumbnailPath)
}

func (s *service) GetAllSeasons() ([]Season, error) {
	rows, err := s.db.Query("SELECT * FROM season")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	seasons := []Season{}
	for rows.Next() {
		var season Season
		if err = rows.Scan(&season.Id, &season.Data.Name, &season.Data.ThumbnailPath, &season.Data.SerieId); err != nil {
			return nil, err
		}
		seasons = append(seasons, season)
	}
	return seasons, err
}
