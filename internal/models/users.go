package models

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// REGISTER USER DANDO ERRO EM DB EXEC -> FAILED TO REGISTER USER
func RegisterUser(username, password string, db *sql.DB) error {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
	if err != sql.ErrNoRows {
		if err == nil {
			return errors.New("username already exists")
		}
		return err
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, string(hashedpassword))
	if err != nil {
		return errors.New("failed to register user")
	}
	return nil
}

func LoginUser(username, password string, db *sql.DB) (int, error) {
	var id int
	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("user does not exist")
		}
		return 0, err
	}

	var storedpassword string
	err = db.QueryRow("SELECT password FROM users WHERE id = ?", id).Scan(&storedpassword)
	if err != nil {
		return 0, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(storedpassword), []byte(password)); err != nil {
		return 0, errors.New("invalid password")
	}
	return id, nil
}
