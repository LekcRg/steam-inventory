package response

import (
	"encoding/json"
	"net/http"

	"github.com/LekcRg/steam-inventory/internal/models"
	"go.uber.org/zap"
)

type Responder struct {
	log *zap.Logger
}

func New(log *zap.Logger) *Responder {
	return &Responder{log: log}
}

func (r *Responder) JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res, err := json.Marshal(data)
	if err != nil {
		r.log.Error("Failed to marshal JSON", zap.Error(err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

		return
	}

	if _, err := w.Write(res); err != nil {
		r.log.Error("Failed to write JSON response", zap.Error(err))
	}
}

func (r *Responder) Message(w http.ResponseWriter, status int, message string) {
	r.JSON(w, status, models.ResponseMessage{
		Message: message,
	})
}

func (r *Responder) Error(w http.ResponseWriter, status int, message string) {
	r.JSON(w, status, models.ResponseError{
		Error: message,
	})
}

func (r *Responder) InternalError(w http.ResponseWriter) {
	r.Error(w, http.StatusInternalServerError, "Internal server error")
}
