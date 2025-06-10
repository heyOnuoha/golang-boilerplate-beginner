package repositories

import (
	"context"
	"net/http"
	"todo-api/database"
	"todo-api/internal/dtos"
	"todo-api/internal/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewTodoRepository(logger *zap.Logger) *TodoRepository {
	return &TodoRepository{
		DB:     database.GetDB(),
		Logger: logger,
	}
}
func (r *TodoRepository) GetTodoItems(ctx context.Context) (dtos.StructuredResponse, error) {
	var todoItems []models.TodoItem

	r.Logger.Info("GetTodoItems request received")

	// Use Preload to load the related Notes for each TodoItem
	if err := r.DB.Preload("Notes").Preload("User").Find(&todoItems).Error; err != nil {
		r.Logger.Error("Failed to retrieve todo items", zap.Error(err))
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: "Failed to retrieve todo items",
			Payload: nil,
		}, err
	}

	return dtos.StructuredResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Todo items retrieved successfully",
		Payload: todoItems,
	}, nil
}

func (r *TodoRepository) CreateTodoItem(ctx context.Context, todoItemDto dtos.CreateTodoItemDto) (dtos.StructuredResponse, error) {
	// Convert DTO to model
	todoItem := models.TodoItem{
		Title:       todoItemDto.Title,
		Description: todoItemDto.Description,
		IsCompleted: false,
	}

	if err := r.DB.Create(&todoItem).Error; err != nil {
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		}, err
	}

	return dtos.StructuredResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Todo item created successfully",
		Payload: todoItem,
	}, nil
}

func (r *TodoRepository) CreateTodoNote(ctx context.Context, todoNoteDto dtos.CreateTodoNoteDto) (dtos.StructuredResponse, error) {

	// Convert DTO to model
	todoNote := models.TodoNote{
		TodoItemID: todoNoteDto.TodoItemID,
		Note:       todoNoteDto.Note,
	}

	if err := r.DB.Create(&todoNote).Error; err != nil {
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		}, err
	}

	return dtos.StructuredResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Todo note created successfully",
		Payload: todoNote,
	}, nil
}

func (r *TodoRepository) UpdateTodoItem(ctx context.Context, todoItemDto dtos.UpdateTodoItemDto) (dtos.StructuredResponse, error) {

	var todoItem models.TodoItem

	if err := r.DB.First(&todoItem, todoItemDto.ID).Error; err != nil {
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "Todo item not found",
			Payload: nil,
		}, nil
	}

	todoItem.Title = todoItemDto.Title
	todoItem.Description = todoItemDto.Description
	todoItem.IsCompleted = todoItemDto.IsCompleted

	if err := r.DB.Save(&todoItem).Error; err != nil {
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		}, err
	}

	return dtos.StructuredResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Todo item updated successfully",
		Payload: todoItem,
	}, nil
}

func (r *TodoRepository) DeleteTodoItem(ctx context.Context, todoItemDto dtos.DeleteTodoItemDto) (dtos.StructuredResponse, error) {
	var todoItem models.TodoItem

	if err := r.DB.First(&todoItem, todoItemDto.ID).Error; err != nil {
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusNotFound,
			Message: "Todo item not found",
			Payload: nil,
		}, nil
	}

	if err := r.DB.Delete(&todoItem).Error; err != nil {
		return dtos.StructuredResponse{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Payload: nil,
		}, err
	}

	return dtos.StructuredResponse{
		Success: true,
		Status:  http.StatusOK,
		Message: "Todo item deleted successfully",
		Payload: nil,
	}, nil
}
