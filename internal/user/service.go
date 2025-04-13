package user

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/pablobdss/Backend-Schedule/internal/auth"
	"github.com/pablobdss/Backend-Schedule/pkg/models"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func RegisterUser(name, email, password string) (*User, error) {

	if name == "" {
		return nil, errors.New("name is empty")
	}

	if len(name) > 0 && regexp.MustCompile(`^\s+$`).MatchString(name) {
		return nil, errors.New("name cannot be just whitespace")
	}

	if email == "" {
		return nil, errors.New("email is empty")
	}

	if len(email) > 0 && regexp.MustCompile(`^\s+$`).MatchString(email) {
		return nil, errors.New("email cannot be just whitespace")
	}

	if !emailRegex.MatchString(email) {
		return nil, errors.New("invalid email format")
	}

	if password == "" {
		return nil, errors.New("password is empty")
	}

	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	if len(password) > 0 && regexp.MustCompile(`^\s+$`).MatchString(password) {
		return nil, errors.New("password cannot be just whitespace")
	}

	HashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: HashedPassword,
	}, nil
}

func LoginUser(db *sql.DB, email, password string) (*models.LoginResponse, error) {
	u, err := FindByEmail(db, email)
	if err != nil {
		return nil, err
	}

	if !auth.VerifyPassword(password, u.Password) {
		return nil, errors.New("invalid credentials")
	}

	return &models.LoginResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}
