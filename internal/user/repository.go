package user

import (
	"database/sql"
	"errors"
	"strings"
)

func CreateUser(db *sql.DB, user *User) error {
	if user == nil {
		return errors.New("user cannot be empty")
	}

	stmt, err := db.Prepare("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if err := stmt.QueryRow(user.Name, user.Email, user.Password).Scan(&user.ID); err != nil {
		return err
	}

	return nil
}

func FindByEmail(db *sql.DB, email string) (*User, error) {

	email = strings.TrimSpace(email)
	if email == "" { 
		return nil, errors.New("email cannot be empty") 
	}

	stmt, err := db.Prepare("SELECT id, name, email, password FROM users WHERE email = $1")
	if err != nil { return nil, err }
	defer stmt.Close()

	var u User

	row := stmt.QueryRow(email)
    err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
    if err == sql.ErrNoRows {
        return nil, errors.New("user not found")
    }
    if err != nil {
        return nil, err
    }


	return &u, nil
}