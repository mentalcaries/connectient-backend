package routes

import (
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleReadiness)
	return mux
}

func handleReadiness(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method, req.URL.Path)
	respondWithJSON(w, http.StatusOK, Response{Message: "Welcome to Connectient"})
}
