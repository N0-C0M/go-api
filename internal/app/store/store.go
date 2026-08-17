package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	// Add fields for your store here, e.g., database connection, etc.
	config *Config
	db     *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	// Implement any necessary cleanup logic here
	s.db.Close()
}
