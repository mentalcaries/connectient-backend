package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mentalcaries/connectient-backend/internal/routes"
)

type apiConfig struct {
	platform  string
	jwtSecret string
}

func main() {

	const port = 4000

	router := routes.NewRouter()
	addr := fmt.Sprintf(":%v", port)
	server := http.Server{
		Addr:    addr,
		Handler: router,
	}
	log.Printf("Server listening on port %v\n", port)
	log.Fatal(server.ListenAndServe())

}
