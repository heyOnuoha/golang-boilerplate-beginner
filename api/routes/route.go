package routes

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

func SetupRoutes(router *mux.Router, logger *zap.Logger) {
	// Create API v1 subrouter
	api := router.PathPrefix("/api/v1").Subrouter()

	// Create todo subrouter and register routes
	todoRouter := api.PathPrefix("/todo").Subrouter()
	HandleTodoRoutes(todoRouter, logger)

	// Create auth subrouter and register routes
	authRouter := api.PathPrefix("/auth").Subrouter()
	HandleAuthRoutes(authRouter, logger)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))
}
