package handlers

import (
	"encoding/json"
	"net/http"
	"todo-api/internal/dtos"
	"todo-api/internal/services"
	"todo-api/internal/utils"

	"go.uber.org/zap"
)

type TodoHandler struct {
	BaseHandler
	service *services.TodoService
}

func NewTodoHandler(logger *zap.Logger) *TodoHandler {
	return &TodoHandler{
		BaseHandler: BaseHandler{
			Logger: logger,
		},
		service: services.NewTodoService(logger),
	}
}

// @Summary Get all Todo Items
// @Description Get all Todo Items from the database
// @Tags todo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} dtos.StructuredResponse "Todo items retrieved successfully"
// @Failure 401 {object} dtos.StructuredResponse "Unauthorized"
// @Failure 500 {object} dtos.StructuredResponse "Internal server error"
// @Router /todo/get-todos [get]
func (h *TodoHandler) GetTodoItems(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("GetTodoItems request received")

	response, err := h.service.GetTodoItems(r.Context())

	if err != nil {
		h.Logger.Error("Failed to get todo items", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}

	// The repository now guarantees that response. Payload will be a slice (possibly empty)
	h.Logger.Info("Todo items retrieved successfully")

	// Return the actual response, not a hardcoded message
	h.ReturnJSONResponse(w, response)
}

// @Summary Create a new Todo Item
// @Description Create a new Todo Item with the provided details
// @Tags todo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param todo body dtos.CreateTodoItemDto true "Todo item data"
// @Success 200 {object} dtos.StructuredResponse "Todo item created successfully"
// @Failure 401 {object} dtos.StructuredResponse "Unauthorized"
// @Failure 500 {object} dtos.StructuredResponse "Internal server error"
// @Router /todo/create-todo-item [post]
func (h *TodoHandler) CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("CreateTodoItem request received")

	var req dtos.CreateTodoItemDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("Failed to decode request body", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}
	defer r.Body.Close()

	userID, err := utils.GetUserIDFromContext(r.Context())
	if err != nil {
		h.Logger.Error("Failed to get user ID from context", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}

	req.UserID = userID

	h.Logger.Debug("Creating todo item", zap.String("title", req.Title))
	response, err := h.service.CreateTodoItem(r.Context(), req)

	if err != nil {
		h.Logger.Error("Failed to create todo item", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}

	h.Logger.Info("Todo item created successfully")
	h.ReturnJSONResponse(w, response)
}

// @Summary Create a new Todo Note
// @Description Create a new Todo Note for an existing Todo Item
// @Tags todo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param todo body dtos.CreateTodoNoteDto true "Todo note data"
// @Success 200 {object} dtos.StructuredResponse "Todo note created successfully"
// @Failure 401 {object} dtos.StructuredResponse "Unauthorized"
// @Failure 500 {object} dtos.StructuredResponse "Internal server error"
// @Router /todo/create-todo-note [post]
func (h *TodoHandler) CreateTodoNote(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("CreateTodoNote request received")

	var req dtos.CreateTodoNoteDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("Failed to decode request body", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}
	defer r.Body.Close()

	h.Logger.Debug("Creating todo note", zap.Uint("todoItemId", req.TodoItemID))
	response, err := h.service.CreateTodoNote(r.Context(), req)

	if err != nil {
		h.Logger.Error("Failed to create todo note", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}

	h.Logger.Info("Todo note created successfully")
	h.ReturnJSONResponse(w, response)
}

// @Summary Update an existing Todo Item
// @Description Update an existing Todo Item with the provided details
// @Tags todo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param todo body dtos.UpdateTodoItemDto true "Todo item update data"
// @Success 200 {object} dtos.StructuredResponse "Todo item updated successfully"
// @Failure 401 {object} dtos.StructuredResponse "Unauthorized"
// @Failure 500 {object} dtos.StructuredResponse "Internal server error"
// @Router /todo/update-todo-item [put]
func (h *TodoHandler) UpdateTodoItem(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("UpdateTodoItem request received")

	var req dtos.UpdateTodoItemDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("Failed to decode request body", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}
	defer r.Body.Close()

	userID, err := utils.GetUserIDFromContext(r.Context())

	if err != nil {
		h.Logger.Error("Failed to get user ID from context", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}

	req.UserID = userID

	h.Logger.Debug("Updating todo item", zap.Uint("id", req.ID))
	response, err := h.service.UpdateTodoItem(r.Context(), req)

	if err != nil {
		h.Logger.Error("Failed to update todo item", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}

	h.Logger.Info("Todo item updated successfully")
	h.ReturnJSONResponse(w, response)
}

// @Summary Delete an existing Todo Item
// @Description Delete an existing Todo Item by ID
// @Tags todo
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param todo body dtos.DeleteTodoItemDto true "Todo item deletion data"
// @Success 200 {object} dtos.StructuredResponse "Todo item deleted successfully"
// @Failure 401 {object} dtos.StructuredResponse "Unauthorized"
// @Failure 500 {object} dtos.StructuredResponse "Internal server error"
// @Router /todo/delete-todo-item [delete]
func (h *TodoHandler) DeleteTodoItem(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("DeleteTodoItem request received")

	var req dtos.DeleteTodoItemDto

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.Error("Failed to decode request body", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}
	defer r.Body.Close()

	h.Logger.Debug("Deleting todo item", zap.Uint("id", req.ID))
	response, err := h.service.DeleteTodoItem(r.Context(), req)

	if err != nil {
		h.Logger.Error("Failed to delete todo item", zap.Error(err))
		h.ReturnJSONResponse(w, dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		})
		return
	}

	h.Logger.Info("Todo item deleted successfully")
	h.ReturnJSONResponse(w, response)
}
