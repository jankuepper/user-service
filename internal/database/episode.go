package database

import "database/sql"

type EpisodeId = int
type EpisodeData struct {
	Name          string
	EpisodePath   string
	SeasonId      SeasonId
	ThumbnailPath string
}

type Episode struct {
	Id   EpisodeId
	Data EpisodeData
}

func (s *service) CreateEpisode(data EpisodeData) (sql.Result, error) {
	const query = `INSERT INTO episode (name, episodepath, thumbnailpath, seasonid) VALUES ($name, $episodepath, $thumbnailpath, $seasonid)`
	statement, _ := s.db.Prepare(query)
	return statement.Exec(data.Name, data.EpisodePath, data.ThumbnailPath, data.SeasonId)
}

func (s *service) GetAllEpisodes() ([]Episode, error) {
	rows, err := s.db.Query("SELECT * FROM episode")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	episodes := []Episode{}
	for rows.Next() {
		var episode Episode
		if err = rows.Scan(&episode.Id, &episode.Data.Name, &episode.Data.EpisodePath, &episode.Data.ThumbnailPath, &episode.Data.SeasonId); err != nil {
			return nil, err
		}
		episodes = append(episodes, episode)
	}
	return episodes, err
}
