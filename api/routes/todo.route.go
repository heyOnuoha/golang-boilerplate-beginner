package routes

import (
	"net/http"
	"todo-api/api/handlers"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func HandleTodoRoutes(api *mux.Router, logger *zap.Logger) {
	// Debug route to confirm this handler is being registered
	logger.Info("Todo routes registered")

	// Root todo endpoint
	todoHandler := handlers.NewTodoHandler(logger)

	// Public routes (if any)
	// No public routes for now

	// Protected routes (require authentication)
	protectedRouter := ApplyAuthMiddleware(api, logger)
	protectedRouter.HandleFunc("/get-todos", todoHandler.GetTodoItems).Methods(http.MethodGet)
	protectedRouter.HandleFunc("/create-todo-item", todoHandler.CreateTodoItem).Methods(http.MethodPost)
	protectedRouter.HandleFunc("/create-todo-note", todoHandler.CreateTodoNote).Methods(http.MethodPost)
	protectedRouter.HandleFunc("/update-todo-item", todoHandler.UpdateTodoItem).Methods(http.MethodPut)
	protectedRouter.HandleFunc("/delete-todo-item", todoHandler.DeleteTodoItem).Methods(http.MethodDelete)
}
