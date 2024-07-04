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

type CreateChirpPayload struct {
	Body string `json:"body"`
}

type UserStore interface {
	CreateUser(user User) error
	GetUsers() ([]User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	GenerateUserID() (int, error)
	UpdateUser(userID int, NewEmail, NewPwHash string) error
}

type User struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	PwHash string `json:"passwordHash"`
}

type CreateUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=130"`
}

type CreateUserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type LoginUserPayload struct {
	Email            string `json:"email" validate:"required,email"`
	Password         string `json:"password" validate:"required"`
	ExpiresInSeconds int    `json:"expires_in_seconds" validate:"number"`
}

type LoginUserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}
