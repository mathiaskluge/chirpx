package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mathiaskluge/chirpx/types"
)

func TestUserServiceHandler(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("fails if user payload is invalid", func(t *testing.T) {
		payload, _ := json.Marshal(
			types.CreateUserPayload{
				Email:    "invalid",
				Password: "123",
			})

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/users", handler.handlerCreateUser)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected code: %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("created user with valid payload", func(t *testing.T) {
		payload, _ := json.Marshal(
			types.CreateUserPayload{
				Email:    "valid2@gmail.com",
				Password: "123abcdefg!@#$*",
			})

		req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/users", handler.handlerCreateUser)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected code: %d, got %d", http.StatusCreated, rr.Code)
		}
	})

}

type mockUserStore struct {
}

func (m *mockUserStore) UpdateUser(userID int, NewEmail, NewPwHash string) error {
	return nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}

func (m *mockUserStore) GetUsers() ([]types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("User not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GenerateUserID() (int, error) {
	return 0, nil
}
