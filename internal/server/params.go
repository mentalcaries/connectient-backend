package server

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func parseId(r *http.Request, param string) (uuid.UUID, error) {
	requestParam := r.PathValue(param)
	if requestParam == "" {
		return uuid.UUID{}, errors.New("missing or invalid ID")
	}
	id, err := uuid.Parse(requestParam)
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}