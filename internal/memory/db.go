package memory

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	db *sql.DB
}

func NewStore() (*Store, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	mokaDir := filepath.Join(home, ".moka")
	if err := os.MkdirAll(mokaDir, 0755); err != nil {
		return nil, err
	}
	dbPath := filepath.Join(mokaDir, "state.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	query := `
		CREATE TABLE IF NOT EXISTS memories (
			id INTEGER PRIMARY KEY autoincrement,
			content TEXT NOT NULL,
			category TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);

		CREATE TABLE IF NOT EXISTS messages (
			id INTEGER PRIMARY KEY autoincrement,
			role TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) AddMemory(content, category string) error {
	_, err := s.db.Exec("INSERT INTO memories (content, category) VALUES (?, ?)", content, category)
	return err
}

func (s *Store) GetRelevantMemories() (string, error) {
	rows, err := s.db.Query("SELECT content FROM memories ORDER BY created_at DESC LIMIT 5")
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var context string
	for rows.Next() {
		var c string
		if err := rows.Scan(&c); err != nil {
			context += "- " + c + "\n"
		}
	}
	return context, nil
}
