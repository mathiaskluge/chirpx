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

func (s *Store) DeleteChirp(chirpID int) error {
	dat, err := s.db.LoadDB()
	if err != nil {
		return fmt.Errorf("DeleteChirp: Failed -> %w", err)
	}

	if dat.Chirps == nil {
		return errors.New("DeleteChirp: No chirps in database")
	}

	_, ok := dat.Chirps[chirpID]
	if !ok {
		return fmt.Errorf("DeleteChirp: Chirp with ID: %v does not exist", chirpID)
	}

	delete(dat.Chirps, chirpID)

	err = s.db.WriteDB(dat)
	if err != nil {
		return fmt.Errorf("DeleteChirp: Failed -> %w", err)
	}

	return nil
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

func (s *Store) GetChirpByID(id int) (types.Chirp, error) {
	dat, err := s.db.LoadDB()
	if err != nil {
		return types.Chirp{}, fmt.Errorf("GetChirpByID: Failed -> %w", err)
	}

	if dat.Chirps == nil {
		return types.Chirp{}, errors.New("GetUChirpByID: No chirps in database")
	}

	chirp, ok := dat.Chirps[id]
	if !ok {
		return types.Chirp{}, fmt.Errorf("GetChirpByID: Chirp with ID: %v does not exist", id)
	}
	return chirp, nil
}

func (s *Store) GetChirpsByAuthor(authorID int, sortOrder string) ([]types.Chirp, error) {
	allChirps, err := s.GetChirps(sortOrder)
	if err != nil {
		return []types.Chirp{}, fmt.Errorf("GetChirpsByID: Failed -> %w", err)
	}

	chirpsWithAuthorID := []types.Chirp{}
	for _, chirp := range allChirps {
		if chirp.AuthorID == authorID {
			chirpsWithAuthorID = append(chirpsWithAuthorID, chirp)
		}
	}

	return chirpsWithAuthorID, nil
}

// Generates a new user ID
// Uses store.GetUsers() to determine next ID
func (s *Store) GenerateChirpID() (int, error) {
	chirps, err := s.GetChirps("asc")
	if err != nil {
		return 0, fmt.Errorf("GenerateChirpID: Failed -> %w", err)
	}
	if len(chirps) == 0 {
		return 1, nil
	}

	return len(chirps) + 1, nil
}

// GetChirps retrieves chirps from the store and sorts them based on sortOrder.
// If sortOrder is "asc", chirps are sorted in ascending order of ID.
// If sortOrder is "desc", chirps are sorted in descending order of ID.
// If sortOrder is empty or any other value, chirps are sorted in ascending order by default.
func (s *Store) GetChirps(sortOrder string) ([]types.Chirp, error) {
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

	switch sortOrder {
	case "desc":
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].ID > chirps[j].ID
		})
	case "asc", "":
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].ID < chirps[j].ID
		})
	default:
		return []types.Chirp{}, fmt.Errorf("getChirps: unsupported sortOrder: %s", sortOrder)
	}

	return chirps, nil
}
