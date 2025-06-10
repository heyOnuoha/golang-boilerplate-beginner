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

func (h *BaseHandler) ReturnJSONResponse(w http.ResponseWriter, response dtos.StructuredResponse) {
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
