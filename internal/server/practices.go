package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	db "github.com/mentalcaries/connectient-backend/internal/database"
)

type Practice struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	ModifiedAt    time.Time  `json:"modified_at"`
	Name          string     `json:"name"`
	City          string     `json:"city"`
	Phone         string     `json:"phone"`
	Email         string     `json:"email"`
	Owner         *uuid.UUID `json:"owner,omitempty"`
	PracticeCode  string     `json:"practice_code"`
	Logo          string     `json:"logo,omitempty"`
	StreetAddress string     `json:"street_address"`
	Instagram     string     `json:"instagram,omitempty"`
	Facebook      string     `json:"facebook,omitempty"`
	Website       string     `json:"website,omitempty"`
}

type PracticeRequest struct {
	ID            *uuid.UUID `json:"id,omitempty"`
	Name          string     `json:"name"`
	City          string     `json:"city"`
	Phone         string     `json:"phone"`
	Email         string     `json:"email"`
	Owner         *uuid.UUID `json:"owner,omitempty"`
	PracticeCode  string     `json:"practice_code"`
	Logo          string     `json:"logo,omitempty"`
	StreetAddress string     `json:"street_address"`
	Instagram     string     `json:"instagram,omitempty"`
	Facebook      string     `json:"facebook,omitempty"`
	Website       string     `json:"website,omitempty"`
}

func (s *Server) handlerPracticeGetAll(w http.ResponseWriter, r *http.Request) {
	practicesDb, err := s.DB.GetPractices(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to fetch practices", err)
		return
	}

	allPractices := []Practice{}
	for _, practice := range practicesDb {
		allPractices = append(allPractices, Practice{
			ID:            practice.ID,
			CreatedAt:     practice.CreatedAt,
			Name:          practice.Name,
			City:          practice.City,
			Phone:         practice.Phone,
			Email:         practice.Email,
			Owner:         practice.Owner,
			PracticeCode:  *practice.PracticeCode,
			Logo:          *practice.Logo,
			StreetAddress: *practice.StreetAddress,
			Instagram:     *practice.Instagram,
			Facebook:      *practice.Facebook,
			Website:       *practice.Website,
		})
	}

	respondWithJSON(w, http.StatusOK, allPractices)
}

func (s *Server) handlerPracticeCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	params := PracticeRequest{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode request", err)
		return
	}

	newPractice, err := s.DB.CreatePractice(r.Context(), db.CreatePracticeParams{
		Name:          params.Name,
		City:          params.City,
		Phone:         params.Phone,
		Email:         params.Email,
		PracticeCode:  &params.PracticeCode,
		Logo:          &params.Logo,
		StreetAddress: &params.StreetAddress,
		Instagram:     &params.Instagram,
		Facebook:      &params.Facebook,
		Website:       &params.Website,
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				errMsg := ""
				switch pgErr.ConstraintName {
				case "practices_email_key":
					errMsg = "Email already in use"
				case "practices_practice_code_key":
					errMsg = "Practice code already taken"
				}
				respondWithError(w, http.StatusConflict, errMsg, err)
				return
			}
		}
		respondWithError(w, http.StatusInternalServerError, "Could not create practice", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, newPractice)
}

func (s *Server) handlerPracticeUpdate(w http.ResponseWriter, r *http.Request) {
	pracId, err := parseId(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Missing or invalid ID", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := PracticeRequest{}

	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode request", err)
		return
	}

	updatedPractice, err := s.DB.UpdatePractice(r.Context(), db.UpdatePracticeParams{
		ID:            pracId,
		Name:          params.Name,
		City:          params.City,
		Phone:         params.Phone,
		Email:         params.Email,
		PracticeCode:  &params.PracticeCode,
		Logo:          &params.Logo,
		StreetAddress: &params.StreetAddress,
		Instagram:     &params.Instagram,
		Facebook:      &params.Facebook,
		Website:       &params.Website,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Invalid Practice ID", err)
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Could not update practice", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Practice{
		ID:            updatedPractice.ID,
		Name:          updatedPractice.Name,
		City:          updatedPractice.City,
		Phone:         updatedPractice.Phone,
		Email:         updatedPractice.Email,
		PracticeCode:  *updatedPractice.PracticeCode,
		Logo:          *updatedPractice.Logo,
		StreetAddress: *updatedPractice.StreetAddress,
		Instagram:     *updatedPractice.Instagram,
		Facebook:      *updatedPractice.Facebook,
		Website:       *updatedPractice.Website,
	})

}
