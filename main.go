package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/highdolen/recommendation_system/db"
	"github.com/highdolen/recommendation_system/internal/api"
)

func main() {
	// Подключаемся к базе
	database := db.Connect()
	defer database.Close()

	// Регистрируем маршруты API
	api.RegisterRoutes(database)

	// Обслуживаем статику и корень
	fs := http.FileServer(http.Dir("./web"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./web/auth.html")
			return
		}
		fs.ServeHTTP(w, r)
	})

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
