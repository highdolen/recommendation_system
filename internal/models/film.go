package models

import (
	"database/sql"
	"fmt"
	"log"
)

type Film struct {
	ID    int64  `json:"id,omitempty"`
	Title string `json:"title"`
	Genre string `json:"genre"`
	Year  int    `json:"year"`
}

func AddFilm(db *sql.DB, film Film) (int64, error) {
	query := "INSERT INTO films (title, genre, year) VALUES ($1, $2, $3) RETURNING film_id"
	var filmID int64
	err := db.QueryRow(query, film.Title, film.Genre, film.Year).Scan(&filmID)
	if err != nil {
		return 0, fmt.Errorf("AddFilm: %w", err)
	}
	return filmID, nil
}

func GetFilmsByGenre(db *sql.DB, genre string) ([]Film, error) {
	query := "SELECT film_id, title, genre, year FROM films WHERE genre = $1"
	rows, err := db.Query(query, genre)
	if err != nil {
		log.Println("GetFilmsByGenre error:", err)
		return nil, fmt.Errorf("GetFilmsByGenre: %w", err)
	}
	defer rows.Close()

	var films []Film
	for rows.Next() {
		var f Film
		if err := rows.Scan(&f.ID, &f.Title, &f.Genre, &f.Year); err != nil {
			return nil, err
		}
		films = append(films, f)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return films, nil
}
