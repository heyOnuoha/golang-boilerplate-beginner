package services

import (
	"context"
	"todo-api/internal/dtos"
	"todo-api/internal/repositories"

	"go.uber.org/zap"
)

type TodoService struct {
	todoRepository *repositories.TodoRepository
}

func NewTodoService(logger *zap.Logger) *TodoService {
	return &TodoService{
		todoRepository: repositories.NewTodoRepository(logger),
	}
}

func (s *TodoService) GetTodoItems(ctx context.Context) (dtos.StructuredResponse, error) {
	response, err := s.todoRepository.GetTodoItems(ctx)
	if err != nil {
		return dtos.StructuredResponse{
			Success: false,
			Status:  500,
			Message: "Failed to retrieve todo items",
			Payload: nil,
		}, err
	}
	return response, nil
}

func (s *TodoService) CreateTodoItem(ctx context.Context, todoItem dtos.CreateTodoItemDto) (dtos.StructuredResponse, error) {

	return s.todoRepository.CreateTodoItem(ctx, todoItem)
}

func (s *TodoService) CreateTodoNote(ctx context.Context, todoNoteDto dtos.CreateTodoNoteDto) (dtos.StructuredResponse, error) {
	return s.todoRepository.CreateTodoNote(ctx, todoNoteDto)
}

func (s *TodoService) UpdateTodoItem(ctx context.Context, todoItemDto dtos.UpdateTodoItemDto) (dtos.StructuredResponse, error) {
	return s.todoRepository.UpdateTodoItem(ctx, todoItemDto)
}

func (s *TodoService) DeleteTodoItem(ctx context.Context, todoItemDto dtos.DeleteTodoItemDto) (dtos.StructuredResponse, error) {
	return s.todoRepository.DeleteTodoItem(ctx, todoItemDto)
}
