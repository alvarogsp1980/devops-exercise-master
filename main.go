package main

import (
	"fmt"
	"net/http"
	"store/pkg/redis"             // Importing custom package for Redis operations.
	httpServer "store/pkg/server" // Alias `httpServer` for custom server package.
	"store/pkg/utils"             // Utility functions package, for example, to read environment variables.

	"github.com/prometheus/client_golang/prometheus/promhttp" // Prometheus library for exposing metrics.
	"github.com/sirupsen/logrus"                              // Logging library.
)

// Global logger variable, using `logrus` for structured, leveled logging.
var logger = logrus.New()

// init function is called when the package is initialized. Used here to configure the logger.
func init() {
	logger.Formatter = &logrus.JSONFormatter{} // Setting log output format to JSON.
	logger.SetLevel(logrus.InfoLevel)          // Setting log level to Info, so debug messages are not logged.
}

func main() {
	logger.Info("Starting application...") // Logging the start of the application.

	// Retrieving Redis address from environment variable, or defaulting to "redis:6379" if not set.
	redisAddr := utils.GetEnv("REDIS_ADDR", "redis:6379")

	// Retrieving application port from environment variable, or defaulting to ":8080".
	port := utils.GetEnv("STORE_PORT", ":8080")

	// Initializing a new Redis repository with the address and logger.
	repo := redis.NewOrdersRepo(redisAddr, logger)
	// Initializing a new HTTP server with the repository and logger.
	server := httpServer.NewServer(repo, logger)
	// Creating a new HTTP ServeMux, which is an HTTP request multiplexer.
	mux := http.NewServeMux()

	// Setting up routes for the server.
	setupRoutes(mux, server)

	// Exposing Prometheus metrics at the `/metrics` endpoint.
	mux.Handle("/metrics", promhttp.Handler())

	// Logging the port on which the server will run.
	fmt.Println("Server is running on port", port)
	// Starting the HTTP server. Logs fatal error if the server fails to start.
	if err := http.ListenAndServe(port, mux); err != nil {
		logger.Fatal("Failed to start server: ", err)
	}
}

// setupRoutes defines routes for the HTTP server and associates them with handler functions.
func setupRoutes(mux *http.ServeMux, server *httpServer.Server) {
	// `/create` endpoint for creating new orders.
	mux.HandleFunc("/create", server.Create)
	// `/order/` endpoint for retrieving orders. It's a prefix route.
	mux.HandleFunc("/order/", server.Get)
}
