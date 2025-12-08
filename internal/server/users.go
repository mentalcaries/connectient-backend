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

type NewUserRequest struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MobilePhone string    `json:"mobile_phone"`
	Email       string    `json:"email"`
	PracticeID  uuid.UUID `json:"practice_id"`
}

type UpdateUserRequest struct {
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	MobilePhone *string `json:"mobile_phone,omitempty"`
	Email       *string `json:"email,omitempty"`
	Role        *string `json:"role,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
	AvatarUrl   *string `json:"avatar_url,omitempty"`
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MobilePhone string    `json:"mobile_phone"`
	Email       string    `json:"email"`
	PracticeID  uuid.UUID `json:"practice_id"`
	Role        string    `json:"role"`
	IsActive    bool      `json:"is_active"`
	AvatarUrl   string    `json:"avatar_url,omitempty"`
}

func (s *Server) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	params := NewUserRequest{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode request", err)
		return
	}
	newUser, err := s.DB.CreateUser(r.Context(), db.CreateUserParams{
		PracticeID:  params.PracticeID,
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		Email:       params.Email,
		MobilePhone: params.MobilePhone,
	})

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				errMsg := ""
				switch pgErr.ConstraintName {
				case "users_email_key":
					errMsg = "Email already in use"
				case "users_mobile_phone_key":
					errMsg = "Mobile number already in use"
				}
				respondWithError(w, http.StatusConflict, errMsg, err)
				return
			}
		}
		respondWithError(w, http.StatusInternalServerError, "Could not create User", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, UserResponse{
		ID:          newUser.ID,
		CreatedAt:   newUser.CreatedAt,
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		MobilePhone: newUser.MobilePhone,
		Email:       newUser.Email,
		PracticeID:  newUser.PracticeID,
		Role:        string(newUser.Role),
		IsActive:    newUser.IsActive,
		AvatarUrl:   newUser.AvatarUrl,
	})
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := parseId(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Missing or invalid ID", err)
		return
	}

	user, err := s.DB.GetUser(r.Context(), userId)
	if err != nil || err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "No user found", err)
		return
	}

	respondWithJSON(w, http.StatusOK, UserResponse{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		ModifiedAt:  user.ModifiedAt,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		MobilePhone: user.MobilePhone,
		Email:       user.Email,
		PracticeID:  user.PracticeID,
		Role:        string(user.Role),
		IsActive:    user.IsActive,
		AvatarUrl:   user.AvatarUrl,
	})

}

func (s *Server) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := parseId(r, "id")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Missing or invalid ID", err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	req := UpdateUserRequest{}

	err = decoder.Decode(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not decode request", err)
		return
	}


	updateParams := db.UpdateUserParams{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		MobilePhone: req.MobilePhone,
		IsActive:    req.IsActive,
		AvatarUrl:   req.AvatarUrl,
	}

	// only set Role if the client sent it
	if req.Role != nil {
		updateParams.Role = db.NullUserRole{
			UserRole: db.UserRole(*req.Role), // "admin" | "provider" | "staff"
			Valid:    true,
		}
	}

	updatedUser, err := s.DB.UpdateUser(r.Context(), db.UpdateUserParams{
		ID:          userId,
		FirstName:   updateParams.FirstName,
		LastName:    updateParams.LastName,
		Email:       updateParams.Email,
		MobilePhone: updateParams.MobilePhone,
		Role:        updateParams.Role,
		IsActive:    updateParams.IsActive,
		AvatarUrl:   updateParams.AvatarUrl,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not update user", err)
		return
	}

	respondWithJSON(w, http.StatusOK, UserResponse{
		ID:          updatedUser.ID,
		CreatedAt:   updatedUser.CreatedAt,
		ModifiedAt:  updatedUser.ModifiedAt,
		FirstName:   updatedUser.FirstName,
		LastName:    updatedUser.LastName,
		MobilePhone: updatedUser.MobilePhone,
		Email:       updatedUser.Email,
		PracticeID:  updatedUser.PracticeID,
		Role:        string(updatedUser.Role),
		IsActive:    updatedUser.IsActive,
		AvatarUrl:   updatedUser.AvatarUrl,
	})
}
