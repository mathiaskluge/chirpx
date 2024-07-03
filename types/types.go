package types

type Chirp struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

type ChirpStore interface {
	CreateChirp(chirp Chirp) error
	GetChirps() ([]Chirp, error)
	GetChirpByID(id int) (*Chirp, error)
	GenerateChirpID() (int, error)
}

type UserStore interface {
	CreateUser(user User) error
	GetUsers() ([]User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	GenerateUserID() (int, error)
}

type User struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	PwHash string `json:"passwordHash"`
}

type CreateUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateChirpPayload struct {
	Body string `json:"body"`
}
