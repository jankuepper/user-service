package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
	CreateTable(query string, name string)
	CreateUser(data UserData) (sql.Result, error)
	GetUserByEmail(email string) (User, error)
	CreateSerie(data SerieData) (sql.Result, error)
	GetAllSeries() ([]Serie, error)
	CreateSeason(data SeasonData) (sql.Result, error)
	GetAllSeasons() ([]Season, error)
	CreateEpisode(data EpisodeData) (sql.Result, error)
	GetAllEpisodes() ([]Episode, error)
}

type service struct {
	db *sql.DB
}

var (
	dburl      = os.Getenv("DB_URL")
	dbInstance *service
)

func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	dbInstance = &service{
		db: db,
	}
	dbInstance.init()
	return dbInstance
}

func (s *service) init() {
	const createUserTable = "CREATE TABLE IF NOT EXISTS user (id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL)"
	s.CreateTable(createUserTable, "user")
	const createJwtTable = "CREATE TABLE IF NOT EXISTS jwt (id INTEGER PRIMARY KEY AUTOINCREMENT, jwt TEXT NOT NULL UNIQUE, userid INTEGER, FOREIGN KEY (userid) REFERENCES user(id))"
	s.CreateTable(createJwtTable, "jwt")
	const createSerieTable = "CREATE TABLE IF NOT EXISTS serie (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, thumbnailpath TEXT)"
	s.CreateTable(createSerieTable, "serie")
	const createSeasonTable = "CREATE TABLE IF NOT EXISTS season (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, thumbnailpath TEXT, serieid INTEGER, FOREIGN KEY (serieid) REFERENCES serie(id))"
	s.CreateTable(createSeasonTable, "season")
	const createEpisodeTable = "CREATE TABLE IF NOT EXISTS episode (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL UNIQUE, episodepath TEXT, thumbnailpath TEXT, seasonid INTEGER, FOREIGN KEY (seasonid) REFERENCES season(id))"
	s.CreateTable(createEpisodeTable, "episode")

	// season 1 for testing
	s.CreateSerie(SerieData{Name: "That 70s show", ThumbnailPath: "https://m.media-amazon.com/images/M/MV5BMzFmZWQ2MDYtZjJlNC00ZDg1LTkwOGItZTAxODA2ZTMzNjQxXkEyXkFqcGc@._V1_.jpg"})
	s.CreateSeason(SeasonData{Name: "Season One", ThumbnailPath: "https://m.media-amazon.com/images/M/MV5BMzFmZWQ2MDYtZjJlNC00ZDg1LTkwOGItZTAxODA2ZTMzNjQxXkEyXkFqcGc@._V1_.jpg", SerieId: 1})
	s.CreateEpisode(EpisodeData{Name: "Pilot", EpisodePath: "Season_1_Disk_1/1_Pilot.mp4", SeasonId: 1, ThumbnailPath: "https://m.media-amazon.com/images/M/MV5BMTcyODE2NzM0Nl5BMl5BanBnXkFtZTgwODQwNTU2MjE@._V1_.jpg"})
	s.CreateEpisode(EpisodeData{Name: "Eric's Birthday", EpisodePath: "Season_1_Disk_1/2_Erics_Birthday.mp4", SeasonId: 1, ThumbnailPath: "https://m.media-amazon.com/images/M/MV5BNzg2MzUwMTg0NF5BMl5BanBnXkFtZTgwMTM4MjY2MjE@._V1_.jpg"})
	s.CreateEpisode(EpisodeData{Name: "Streaking", EpisodePath: "Season_1_Disk_1/3_Streaking.mp4", SeasonId: 1, ThumbnailPath: "https://m.media-amazon.com/images/M/MV5BZDQzOGM2NTItZjMyMS00OTg2LTk1N2QtMzY2YmYzZjE0NjdhXkEyXkFqcGc@._V1_.jpg"})
	s.CreateEpisode(EpisodeData{Name: "Battle of the sexiest", EpisodePath: "Season_1_Disk_1/4_Battle_of_the_sexiest.mp4", SeasonId: 1, ThumbnailPath: "https://m.media-amazon.com/images/M/MV5BMjMxNTcwMzExOF5BMl5BanBnXkFtZTgwMjE5ODU2MjE@._V1_.jpg"})
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("%s", fmt.Sprintf("db down: %v", err)) // Log the error and terminate the program
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	return s.db.Close()
}

func (s *service) CreateTable(query string, name string) {
	statement, err := s.db.Prepare(query)
	if err != nil {
		log.Printf(`Error in creating table %s`, name)
	} else {
		log.Printf("Successfully created table %s!", name)
	}
	statement.Exec()
}
