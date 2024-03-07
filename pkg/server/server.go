// Package server defines the HTTP server and its routes for creating and retrieving orders.
package server

import (
	"encoding/json"    // For encoding and decoding JSON data.
	"net/http"         // To use the HTTP server functionalities.
	"store/pkg/domain" // Importing the domain package for order-related structures.
	"store/pkg/utils"  // For utility functions, like generating new IDs.
	"strings"          // To manipulate strings, specifically to parse the URL.

	"github.com/sirupsen/logrus" // Importing logrus for logging.
)

// Server struct contains the dependencies needed for the server to operate.
type Server struct {
	repo   domain.OrdersRepo // Interface for order repository, abstracts the storage layer.
	logger *logrus.Logger    // Logger for logging messages, errors, etc.
}

// NewServer initializes a new Server instance with the given orders repository and logger.
func NewServer(repo domain.OrdersRepo, logger *logrus.Logger) *Server {
	return &Server{
		repo:   repo,   // Order repository for data storage operations.
		logger: logger, // Logger instance for logging.
	}
}

// SetupRoutes defines the HTTP routes the server will respond to.
func (s *Server) SetupRoutes() http.Handler {
	mux := http.NewServeMux()           // Creates a new HTTP serve mux for routing.
	mux.HandleFunc("/create", s.Create) // Route for creating a new order.
	mux.HandleFunc("/order/", s.Get)    // Route for retrieving an existing order by ID.
	return mux                          // Returns the configured router.
}

// Create handles HTTP requests for creating a new order.
func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Gets the context from the HTTP request.
	order := domain.Order{
		ID:     utils.NewID(), // Generates a new unique ID for the order.
		Status: "CREATED",     // Sets the initial status of the order.
	}
	err := s.repo.Save(ctx, order) // Saves the new order to the repository.
	if err != nil {
		// Logs and returns an error if the order cannot be saved.
		s.logger.WithFields(logrus.Fields{"orderID": order.ID, "error": err.Error()}).Error("Error saving the order")
		w.WriteHeader(http.StatusInternalServerError) // Sends an HTTP 500 response.
		return
	}
	// Logs the successful creation and returns the order ID.
	s.logger.WithFields(logrus.Fields{"orderID": order.ID}).Info("Order successfully created")
	w.WriteHeader(http.StatusCreated) // Sends an HTTP 201 response.
	w.Write([]byte(order.ID))         // Writes the order ID in the response body.
}

// Get handles HTTP requests for retrieving an existing order by its ID.
func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context() // Gets the context from the HTTP request.
	var id string
	parts := strings.Split(r.URL.String(), "/") // Splits the URL to extract the order ID.
	if len(parts) > 2 {
		id = parts[2] // The order ID is expected to be the third part of the URL.
	}
	order, err := s.repo.Get(ctx, id) // Retrieves the order from the repository using its ID.
	if err != nil {
		// Logs and returns an error if the order cannot be retrieved.
		s.logger.WithFields(logrus.Fields{"orderID": id, "error": err.Error()}).Error("Error retrieving the order")
		w.WriteHeader(http.StatusInternalServerError) // Sends an HTTP 500 response.
		return
	}
	// Logs the successful retrieval.
	s.logger.WithFields(logrus.Fields{"orderID": id}).Info("Order successfully retrieved")
	w.WriteHeader(http.StatusOK)     // Sends an HTTP 200 response.
	body, err := json.Marshal(order) // Encodes the order as JSON.
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sends an HTTP 500 response if encoding fails.
		return
	}
	w.Write(body) // Writes the JSON-encoded order in the response body.
}
