package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mentalcaries/connectient-backend/internal/database"
)

type Response struct {
	Message string `json:"message"`
}

type Appointment struct {
	ID              uuid.UUID `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	ModifiedAt      time.Time `json:"modified_at"`
	Email           string    `json:"email"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	MobilePhone     string    `json:"mobile_phone"`
	RequestDate     time.Time `json:"requested_date"`
	IsEmergency     bool      `json:"is_emergency"`
	Description     string    `json:"description"`
	AppointmentType string    `json:"appointment_type"`
	IsScheduled     bool      `json:"is_scheduled"`
	ScheduledDate   time.Time `json:"scheduled_date"`
	CreatedBy       uuid.UUID `json:"created_by"`
	ScheduledBy     uuid.UUID `json:"scheduled_by"`
	IsCancelled     bool      `json:"is_cancelled"`
	RequestedTime   string    `json:"requested_time"`
	ScheduledTime   time.Time `json:"scheduled_time"`
	PracticeId      uuid.UUID `json:"practice_id"`
}

type AppointmenRequest struct {
	Email           string         `json:"email"`
	FirstName       string         `json:"first_name"`
	LastName        string         `json:"last_name"`
	MobilePhone     string         `json:"mobile_phone"`
	RequestDate     time.Time      `json:"requested_date"`
	IsEmergency     sql.NullBool   `json:"is_emergency"`
	Description     sql.NullString `json:"description"`
	AppointmentType sql.NullString `json:"appointment_type"`
	IsScheduled     bool           `json:"is_scheduled"`
	ScheduledDate   time.Time      `json:"scheduled_date"`
	CreatedBy       uuid.UUID      `json:"created_by"`
	ScheduledBy     uuid.UUID      `json:"scheduled_by"`
	IsCancelled     bool           `json:"is_cancelled"`
	RequestedTime   string         `json:"requested_time,omitempty"`
	ScheduledTime   time.Time      `json:"scheduled_time"`
	PracticeId      uuid.UUID      `json:"practice_id"`
}

func NewRouter(s *Server) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleReadiness)
	mux.HandleFunc("POST /appointments", s.handlerAppointmentsCreate)
	return mux
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	respondWithJSON(w, http.StatusOK, Response{Message: "Welcome to Connectient"})
}

func (s *Server) handlerAppointmentsCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := AppointmenRequest{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode request", err)
	}

	appointment, err := s.DB.CreateAppointment(r.Context(), database.CreateAppointmentParams{
		FirstName:       params.FirstName,
		LastName:        params.LastName,
		MobilePhone:     params.MobilePhone,
		Email:           params.Email,
		RequestedDate:   params.RequestDate,
		RequestedTime:   params.RequestedTime,
		AppointmentType: params.AppointmentType,
		Description:     params.Description,
		IsEmergency:     params.IsEmergency,
		PracticeID:      params.PracticeId,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not save to database", err)
	}
	respondWithJSON(w, http.StatusCreated, appointment)
}
