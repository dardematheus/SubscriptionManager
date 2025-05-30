package models

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username, password string, db *sql.DB) error {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
	if err != sql.ErrNoRows {
		if err == nil {
			return errors.New("Username already exists")
		}
		return err
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Error hashing password")
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, string(hashedpassword))
	if err != nil {
		return errors.New("Failed to register user")
	}
	return nil
}

func LoginUser(username, password string, db *sql.DB) error {

}
