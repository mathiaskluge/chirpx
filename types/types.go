package types

type Chirp struct {
	ID       int    `json:"id"`
	Body     string `json:"body"`
	AuthorID int    `json:"author_id"`
}

type ChirpStore interface {
	CreateChirp(chirp Chirp) error
	GetChirps() ([]Chirp, error)
	GetChirpByID(id int) (Chirp, error)
	GenerateChirpID() (int, error)
	DeleteChirp(chirpID int) error
}

type CreateChirpPayload struct {
	Body string `json:"body"`
}

type CreateChirpResponse struct {
	ID       int    `json:"id"`
	Body     string `json:"body"`
	AuthorID int    `json:"author_id"`
}

type UserStore interface {
	CreateUser(user User) error
	GetUsers() ([]User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByID(id int) (User, error)
	GenerateUserID() (int, error)
	UpdateUser(userID int, NewEmail, NewPwHash string) error
	CreateSession(token string, userID int, exppiresInSeconds int) error
	GetSession(token string) (Session, error)
	UpdateSession(token string, session Session) error
	UpgradeUser(userID int) error
}

type User struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	PwHash     string `json:"passwordHash"`
	IsUpgraded bool   `json:"is_chirpy_red"`
}

type CreateUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4,max=130"`
}

type CreateUserResponse struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	IsUpgraded bool   `json:"is_chirpy_red"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse struct {
	ID         int    `json:"id"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Session    string `json:"refresh_token"`
	IsUpgraded bool   `json:"is_chirpy_red"`
}

type Session struct {
	ExpiresAt int64  `json:"expires_at"`
	UserID    int    `json:"user_id"`
	Token     string `json:"token"`
	Revoked   bool   `json:"is_revoked"`
}

type UpgradeUserPayload struct {
	Event string `json:"event"`
	Data  UserID `json:"data"`
}

type UserID struct {
	UserID int `json:"user_id"`
}
