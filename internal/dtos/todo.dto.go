package dtos

import "time"

// TodoItemDto represents a todo item in the system
// @Description A todo item with all its details
type TodoItemDto struct {
	// Unique identifier
	// @example 1
	ID uint `json:"id" example:"1"`
	// Title of the todo item
	// @example Buy groceries
	Title string `json:"title" example:"Buy groceries"`
	// Optional description with details
	// @example Milk, eggs, bread, and cheese
	Description string `json:"description" example:"Milk, eggs, bread, and cheese"`
	// Whether the todo item is completed
	// @example false
	IsCompleted bool `json:"isCompleted" example:"false"`
	// When the todo item was created
	// @example 2025-06-10T10:30:00Z
	CreatedAt time.Time `json:"createdAt" example:"2025-06-10T10:30:00Z"`
	// When the todo item was last updated
	// @example 2025-06-10T10:30:00Z
	UpdatedAt time.Time `json:"updatedAt" example:"2025-06-10T10:30:00Z"`
}

// GetTodoItemDto represents the data needed to retrieve a specific todo item
// @Description Data for retrieving a specific todo item by ID
type GetTodoItemDto struct {
	// ID of the todo item to retrieve
	// @example 1
	ID uint `json:"id" example:"1"`
}

// CreateTodoItemDto represents the data needed to create a new todo item
// @Description Data for creating a new todo item
type CreateTodoItemDto struct {
	// Title of the todo item (3-255 characters)
	// @example Buy groceries
	Title string `json:"title" validate:"required;min=3;max=255" example:"Buy groceries"`
	// Optional description with details (max 255 characters)
	// @example Milk, eggs, bread, and cheese
	Description string `json:"description" validate:"max=255" example:"Milk, eggs, bread, and cheese"`

	// User ID associated with the todo item
	// @example 1
	UserID uint `json:"userId" example:"1"`
}

// UpdateTodoItemDto represents the data needed to update an existing todo item
// @Description Data for updating an existing todo item
type UpdateTodoItemDto struct {
	// ID of the todo item to update
	// @example 1
	ID uint `json:"id" example:"1"`
	// Updated title (3-255 characters)
	// @example Buy groceries and household items
	Title string `json:"title" validate:"min=3;max=255" example:"Buy groceries and household items"`
	// Updated description (max 255 characters)
	// @example Milk, eggs, bread, cheese, and cleaning supplies
	Description string `json:"description" validate:"max=255" example:"Milk, eggs, bread, cheese, and cleaning supplies"`
	// Updated completion status
	// @example true
	IsCompleted bool `json:"isCompleted" example:"true"`

	// User ID associated with the todo item
	// @example 1
	UserID uint `json:"userId" example:"1"`
}

// CreateTodoNoteDto represents the data needed to create a note for a todo item
// @Description Data for creating a new note attached to a todo item
type CreateTodoNoteDto struct {
	// ID of the todo item this note belongs to
	// @example 1
	TodoItemID uint `json:"todoItemId" example:"1"`
	// Content of the note
	// @example Don't forget to check expiration dates
	Note string `json:"note" example:"Don't forget to check expiration dates"`
}

// DeleteTodoItemDto represents the data needed to delete an existing todo item
// @Description Data for deleting an existing todo item

type DeleteTodoItemDto struct {
	// ID of the todo item to delete
	// @example 1
	ID uint `json:"id" example:"1"`
}
