package userservice

import "errors"

type UserRepository interface {
	CreateUser(user User) error
	GetUserByEmail(email string) (User, error)
}

type InMemoryUserRepository struct {
	users map[string]User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]User),
	}
}

func (repo *InMemoryUserRepository) CreateUser(user User) error {
	repo.users[user.Email] = user
	return nil
}

func (repo *InMemoryUserRepository) GetUserByEmail(email string) (User, error) {
	user, exists := repo.users[email]
	if !exists {
		return User{}, ErrUserNotFound
	}
	return user, nil
}

var ErrUserNotFound = errors.New("user not found")
