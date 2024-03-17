package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/AlfonsoRamat/go_backend/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("Port is not defined")
	}

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("Connection string must be provided")
	}

	connection, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal("Can't connect to the databaseñ")
	}

	apiCfg := apiConfig{
		DB: database.New(connection),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handleReadiness)
	v1Router.Get("/users", apiCfg.handleCreateUser)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + PORT,
	}

	log.Printf("Server starting on port: %v", PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
