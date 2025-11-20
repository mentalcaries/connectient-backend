package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/mentalcaries/connectient-backend/internal/database"
)

type Server struct {
	port      int
	Platform  string
	JWTSecret string
	DB        *db.Queries
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

	dbConn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbQueries := db.New(dbConn)
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
