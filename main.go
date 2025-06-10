package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"
	"time"
	"todo-api/api/routes"
	"todo-api/config"
	"todo-api/database"
	_ "todo-api/docs"
	"todo-api/internal/logger"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"

	"os/signal"
)

// @title           Go Todo API
// @version         1.0
// @description     A RESTful API for Todo management built with Go and PostgreSQL
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Enter the token with the `Bearer: ` prefix, e.g. 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...'

func main() {
	// Initialize the application
	cfg, error := config.LoadConfig()

	if error != nil {
		panic("failed to load config")
	}

	logger.InitLogger(cfg.Env)
	// Ensure logger syncs before program exits
	defer zap.L().Sync()

	err := database.InitDatabase(&cfg.Database)

	if err != nil {
		panic("failed to connect to database")
	}

	err = database.Migrate()

	if err != nil {
		fmt.Printf("Migration error: %v\n", err)
		panic("failed to migrate database")
	}

	router := mux.NewRouter()

	routes.SetupRoutes(router, zap.L())

	// Root route for basic testing
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Todo API"))
	}).Methods(http.MethodGet)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: corsMiddleware.Handler(router),
	}

	defer database.CloseDB()

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				fmt.Printf("HTTP server error: %v\n", err)
			}
		}
	}()

	fmt.Println("Server started on port " + cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Server shutting down...")

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait until the timeout deadline
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v\n", err)
	}

	fmt.Println("Server exited properly")
}
