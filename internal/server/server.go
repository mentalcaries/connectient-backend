package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/mentalcaries/connectient-backend/internal/database"
)

type Server struct {
	port      int
	Platform  string
	JWTSecret string
	DB        *database.Queries
}

func NewServer() *http.Server {

	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0  {
		log.Fatal("PORT value must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbQueries := database.New(dbConn)
	apiCfg := Server{
		DB:   dbQueries,
		port: port,
	}

	router := NewRouter(&apiCfg)
	addr := fmt.Sprintf(":%v", port)
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}

    return &server
}
