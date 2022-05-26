package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	config   *Config
	db       *sql.DB
	itemRepo *ItemRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DBUrl)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Storage) Close() {
	if err := s.db.Close(); err != nil {
		log.Fatal(err)
	}
}

func (s *Storage) CreateSchema() {
	s.Open()
	defer s.Close()
	q, _ := s.db.Prepare(`
		CREATE TABLE IF NOT EXISTS "items" (
		    "id" SERIAL PRIMARY KEY,
		    "create_at" TIMESTAMP NOT NULL,
		    "jsondata" json NOT NULL
		    )
		`)
	if _, err := q.Exec(); err != nil {
		log.Fatal(err)
	}
	log.Println("> SCHEMA CREATED <")
}

func (s *Storage) Item() *ItemRepository {
	if s.itemRepo != nil {
		return s.itemRepo
	} else {
		s.itemRepo = &ItemRepository{
			storage: s,
		}
	}
	return s.itemRepo
}
