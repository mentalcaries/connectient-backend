package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mentalcaries/connectient-backend/internal/database"
	"github.com/mentalcaries/connectient-backend/internal/routes"
)

type apiConfig struct {
	platform  string
	jwtSecret string
	db        *database.Queries
}

func main() {
	const port = 4000

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbQueries := database.New(dbConn)
	apiCfg := apiConfig{
		db: dbQueries,
	}

	router := routes.NewRouter()
	addr := fmt.Sprintf(":%v", port)
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}
	log.Printf("Server listening on port %v\n", port)
	log.Fatal(server.ListenAndServe())

}
