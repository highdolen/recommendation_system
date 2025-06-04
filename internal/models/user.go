package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

func AddUser(db *sql.DB, user User) (int64, error) {
	query := `INSERT INTO users (name, email, nickname) VALUES ($1, $2, $3) RETURNING user_id`
	var user_id int64
	err := db.QueryRow(query, user.Name, user.Email, user.Nickname).Scan(&user_id)
	if err != nil {
		return 0, fmt.Errorf("AddUser: %w", err)
	}
	return user_id, nil
}

func FindUserByEmailOrUsername(db *sql.DB, identifier string) (*User, error) {
	query := `SELECT user_id, name, email, nickname FROM users WHERE email = $1 OR nickname = $1`
	var user User
	err := db.QueryRow(query, identifier).Scan(&user.ID, &user.Name, &user.Email, &user.Nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // не найден
		}
		return nil, fmt.Errorf("FindUserByEmailOrUsername: %w", err)
	}
	return &user, nil
}
