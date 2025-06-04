package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/highdolen/recommendation_system/internal/models"
)

// Регистрируем маршруты API
func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/models/users", addUserHandler(db))
	http.HandleFunc("/models/films", addFilmHandler(db))
	http.HandleFunc("/models/films/bygenre", getFilmsByGenreHandler(db))
	http.HandleFunc("/api/login", LoginHandler(db))
}

// Добавление пользователя
func addUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Невалидный JSON", http.StatusBadRequest)
			return
		}

		id, err := models.AddUser(db, user)
		if err != nil {
			http.Error(w, "Ошибка при добавлении пользователя: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int64{"user_id": id})
	}
}

// Добавление фильма
func addFilmHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		var film models.Film
		if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
			http.Error(w, "Невалидный JSON", http.StatusBadRequest)
			return
		}

		log.Printf("DEBUG film: %+v\n", film) // <-- теперь лог в нужном месте

		id, err := models.AddFilm(db, film)
		if err != nil {
			log.Println("Ошибка при добавлении фильма:", err)
			http.Error(w, "Ошибка при добавлении фильма: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int64{"film_id": id})
	}
}

// Получение фильмов по жанру
func getFilmsByGenreHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		genre := r.URL.Query().Get("genre")
		if genre == "" {
			http.Error(w, "Параметр genre обязателен", http.StatusBadRequest)
			return
		}

		films, err := models.GetFilmsByGenre(db, genre)
		if err != nil {
			http.Error(w, "Ошибка при получении фильмов: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(films)
	}
}
