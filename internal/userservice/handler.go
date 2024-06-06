package userservice

import (
	"encoding/json"
	"net/http"

	"github.com/shchepelevnikita/todo-user-service/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

var userRepo = NewInMemoryUserRepository()

func Register(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.JSON(w, http.StatusBadRequest, response.Message{Message: "Invalid request payload"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Message{Message: "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := userRepo.CreateUser(user); err != nil {
		response.JSON(w, http.StatusInternalServerError, response.Message{Message: "Error creating user"})
		return
	}
	response.JSON(w, http.StatusCreated, response.Message{Message: "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		response.JSON(w, http.StatusBadRequest, response.Message{Message: "Invalid request payload"})
		return
	}

	user, err := userRepo.GetUserByEmail(credentials.Email)
	if err != nil {
		response.JSON(w, http.StatusUnauthorized, response.Message{Message: "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		response.JSON(w, http.StatusUnauthorized, response.Message{Message: "Invalid email or password"})
		return
	}

	response.JSON(w, http.StatusOK, response.Message{Message: "User logged in successfully"})
}

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/register", Register)
	mux.HandleFunc("/users/login", Login)
	return mux
}
