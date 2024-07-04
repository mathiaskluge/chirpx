package user

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

// creates a new store object with a pointer to the provided db
func NewStore(db *db.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) UpdateUser(userID int, NewEmail, NewPwHash string) error {

	dat, err := s.db.LoadDB()
	if err != nil {
		return fmt.Errorf("UpdateUser: Failed -> %w", err)
	}

	if dat.Users == nil {
		return errors.New("UpdateUser: no users in db")
	}

	curUser, ok := dat.Users[userID]
	if !ok {
		return errors.New("UpdateUser: User does not exist")
	}

	curUser.Email = NewEmail
	curUser.PwHash = NewPwHash
	dat.Users[userID] = curUser

	err = s.db.WriteDB(dat)
	if err != nil {
		return fmt.Errorf("CreateUser: Failed -> %w", err)
	}

	return nil
}

// Creates a new User and writes changes to disk
func (s *Store) CreateUser(user types.User) error {
	dat, err := s.db.LoadDB()
	if err != nil {
		return fmt.Errorf("CreateUser: Failed -> %w", err)
	}

	if dat.Users == nil {
		dat.Users = map[int]types.User{}
	} else {
		_, ok := dat.Users[user.ID]
		if ok {
			return fmt.Errorf("CreateUser: User with ID: %v already exists", user.ID)
		}
	}

	dat.Users[user.ID] = user
	err = s.db.WriteDB(dat)
	if err != nil {
		return fmt.Errorf("CreateUser: Failed -> %w", err)
	}

	return nil
}

// Returns pointer to a user if it exists, error otherwise
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	dat, err := s.db.LoadDB()
	if err != nil {
		return &types.User{}, fmt.Errorf("GetUserbyEmail: Failed -> %w", err)
	}

	userMap := make(map[string]types.User)
	for _, user := range dat.Users {
		userMap[user.Email] = user
	}
	user, ok := userMap[email]
	if !ok {
		return &types.User{}, fmt.Errorf("GetUserbyEmail: User with Email: %s does not exist", email)
	}
	return &user, nil
}

// Returns pointer to a user if it exists, error otherwise
func (s *Store) GetUserByID(id int) (*types.User, error) {
	dat, err := s.db.LoadDB()
	if err != nil {
		return &types.User{}, fmt.Errorf("GetUserbyID: Failed -> %w", err)
	}

	if dat.Users == nil {
		return &types.User{}, errors.New("GetUserbyID: No users in database")
	}

	user, ok := dat.Users[id]
	if !ok {
		return &types.User{}, fmt.Errorf("GetUserbyID: User with ID: %v does not exist", id)
	}
	return &user, nil
}

// Generates a new user ID
// Uses store.GetUsers() to determine next ID
func (s *Store) GenerateUserID() (int, error) {
	users, err := s.GetUsers()
	if err != nil {
		return 0, fmt.Errorf("GenerateUserID: Failed -> %w", err)
	}
	if len(users) == 0 {
		return 1, nil
	}

	return len(users) + 1, nil
}

// Returns a sorted (by id) array of users
func (s *Store) GetUsers() ([]types.User, error) {
	dat, err := s.db.LoadDB()
	if err != nil {
		return []types.User{}, fmt.Errorf("GetUsers: Failed -> %w", err)
	}

	if len(dat.Users) == 0 {
		return []types.User{}, nil
	}

	users := make([]types.User, 0, len(dat.Users))
	for _, u := range dat.Users {
		users = append(users, u)
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	return users, nil
}
