package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mentalcaries/connectient-backend/internal/server"
)

func main() {

	server := server.NewServer()
	fmt.Printf("Server up and running at: http://locahost%v\n", server.Addr)

	log.Fatal(server.ListenAndServe())

}
