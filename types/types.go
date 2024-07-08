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
	CreateSession(token string, userID int, exppiresInSeconds int) error
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
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Token   string `json:"token"`
	Session string `json:"refresh_token"`
}

type Session struct {
	ExpiresAt int64  `json:"expires_at"`
	UserID    int    `json:"user_id"`
	Token     string `json:"token"`
}
