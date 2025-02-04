package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Siddhant6674/ECOM/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserstore{}
	handler := NewHandler(userStore)

	t.Run("Should fail if user payload is invalid", func(t *testing.T) {
		payload := &types.RegisterUserPayload{
			FirstName: "Siddhant",
			LastName:  "Dutal",
			Email:     "siddhant@gmail.com",
			Password:  "12345678",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d got %d", http.StatusBadRequest, rr.Code)
		}

	})

	t.Run("should correctly register the user", func(t *testing.T) {
		payload := &types.RegisterUserPayload{
			FirstName: "Siddhant",
			LastName:  "Dutal",
			Email:     "siddhant@gmail.com",
			Password:  "12345678",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %d got %d", http.StatusCreated, rr.Code)
		}

	})
}

type mockUserstore struct{}

func (m *mockUserstore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}
func (m *mockUserstore) GetUserByID(ID int) (*types.User, error) {
	return nil, nil
}
func (m *mockUserstore) CreateUser(types.User) error {
	return nil
}
