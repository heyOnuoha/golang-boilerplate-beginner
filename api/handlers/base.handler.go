package handlers

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/dtos"

	"go.uber.org/zap"
)

type BaseHandler struct {
	Logger *zap.Logger
}

func (h *BaseHandler) ReturnJSONResponse(w http.ResponseWriter, response dtos.ApiResponse) {
	responseJSON, err := json.Marshal(response)

	if err != nil {
		h.Logger.Error("Failed to marshal response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	w.Write(responseJSON)
}

func (h *BaseHandler) DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		h.Logger.Error("Failed to decode request body", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.ApiResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Payload: nil,
		})
		return false
	}
	defer r.Body.Close()
	return true
}
