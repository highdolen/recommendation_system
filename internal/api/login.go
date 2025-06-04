package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/highdolen/recommendation_system/internal/models"
)

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		var input struct {
			Identifier string `json:"identifier"` // email или username
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil || input.Identifier == "" {
			http.Error(w, "Невалидный ввод", http.StatusBadRequest)
			return
		}

		user, err := models.FindUserByEmailOrUsername(db, input.Identifier)
		if err != nil {
			http.Error(w, "Ошибка при поиске пользователя: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int64{"user_id": user.ID})
	}
}
