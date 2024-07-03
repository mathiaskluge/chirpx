package chirp

import (
	"errors"
	"fmt"
	"sort"

	"github.com/mathiaskluge/chirpx/db"
	"github.com/mathiaskluge/chirpx/types"
)

type Store struct {
	db *db.DB
}

func NewStore(db *db.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateChirp(chirp types.Chirp) error {
	dat, err := s.db.LoadDB()
	if err != nil {
		return fmt.Errorf("CreateChirp: Failed -> %w", err)
	}

	if dat.Chirps == nil {
		dat.Chirps = map[int]types.Chirp{}
	} else {
		_, ok := dat.Chirps[chirp.ID]
		if ok {
			return fmt.Errorf("CreateChirp: Chirp with ID: %v already exists", chirp.ID)
		}
	}

	dat.Chirps[chirp.ID] = chirp
	err = s.db.WriteDB(dat)
	if err != nil {
		return fmt.Errorf("CreateChirp: Failed -> %w", err)
	}

	return nil
}

func (s *Store) GetChirpByID(id int) (*types.Chirp, error) {
	dat, err := s.db.LoadDB()
	if err != nil {
		return &types.Chirp{}, fmt.Errorf("GetChirpByID: Failed -> %w", err)
	}

	if dat.Chirps == nil {
		return &types.Chirp{}, errors.New("GetUChirpByID: No chirps in database")
	}

	chirp, ok := dat.Chirps[id]
	if !ok {
		return &types.Chirp{}, fmt.Errorf("GetChirpByID: Chirp with ID: %v does not exist", id)
	}
	return &chirp, nil
}

// Generates a new user ID
// Uses store.GetUsers() to determine next ID
func (s *Store) GenerateChirpID() (int, error) {
	chirps, err := s.GetChirps()
	if err != nil {
		return 0, fmt.Errorf("GenerateChirpID: Failed -> %w", err)
	}
	if len(chirps) == 0 {
		return 1, nil
	}

	return len(chirps) + 1, nil
}

func (s *Store) GetChirps() ([]types.Chirp, error) {
	dat, err := s.db.LoadDB()
	if err != nil {
		return []types.Chirp{}, fmt.Errorf("GetChirp: Failed -> %w", err)
	}

	if len(dat.Chirps) == 0 {
		return []types.Chirp{}, nil
	}

	chirps := make([]types.Chirp, 0, len(dat.Chirps))
	for _, u := range dat.Chirps {
		chirps = append(chirps, u)
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	return chirps, nil
}
