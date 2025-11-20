package server

import (
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func NewRouter(s *Server) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleReadiness)
	mux.HandleFunc("GET /appointments", s.handlerAppointmentsGetAll)
	mux.HandleFunc("GET /appointments/{id}", s.handlerGetAppointmentById)
	mux.HandleFunc("POST /appointments", s.handlerAppointmentsCreate)
	return mux
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	respondWithJSON(w, http.StatusOK, Response{Message: "Welcome to Connectient"})
}
