package db

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/mathiaskluge/chirpx/types"
)

type DB struct {
	path string
	mu   *sync.Mutex
}

type DBStructure struct {
	Chirps map[int]types.Chirp `json:"chirps"`
	Users  map[int]types.User  `json:"users"`
}

// Reads the database file (db.path) into memory
// Returns the DBStructure and any potential errors
func (db *DB) LoadDB() (DBStructure, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	content, err := os.ReadFile(db.path)
	if err != nil {
		return DBStructure{}, fmt.Errorf("loadDB: failed to read database file: %w", err)
	}

	var dat DBStructure
	if len(content) > 0 {
		if err := json.Unmarshal(content, &dat); err != nil {
			return DBStructure{}, fmt.Errorf("loadDB: failed to unmarshal JSON: %w", err)
		}
	}

	return dat, nil
}

// This function writes the DBStructure to the database file (db.path).
func (db *DB) WriteDB(dbStructure DBStructure) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	dat, err := json.MarshalIndent(dbStructure, "", "    ")
	if err != nil {
		return fmt.Errorf("writeDB: failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(db.path, dat, 0644); err != nil {
		return fmt.Errorf("writeDB: failed to write file: %w", err)
	}

	return nil
}

// Initializes a new DB struct at the configured path.
// It ensures the database file exists or creates it if it doesn't
func NewDB(path string) (*DB, error) {
	newDB := DB{
		path: path,
		mu:   &sync.Mutex{},
	}
	if err := newDB.ensureDB(); err != nil {
		return nil, err
	}
	return &newDB, nil
}

// Ensures that the database file exists.
// It attempts to open the file and creates it if it doesn't exist.
func (db *DB) ensureDB() error {
	if _, err := os.Open(db.path); err != nil {
		_, err := os.Create(db.path)
		if err != nil {
			return fmt.Errorf("ensureDB: db creation failed: %w", err)
		}
	}

	return nil
}
