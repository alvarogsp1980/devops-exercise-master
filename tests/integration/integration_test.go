// Package integration contains integration tests for the application's HTTP endpoints.
package integration

import (
	"encoding/json"     // For encoding and decoding JSON data in responses.
	"fmt"               // For printing formatted output, useful in logging responses.
	"net/http"          // To access HTTP constants and functionalities.
	"net/http/httptest" // For creating test servers and recording responses.
	"testing"           // For writing test cases.

	"github.com/stretchr/testify/assert" // A toolkit with common assertions for testing.
)

// TestAppEndpointsIntegration tests the application's endpoints for creating and retrieving orders.
func TestAppEndpointsIntegration(t *testing.T) {
	server := setupTestServer() // Sets up a test HTTP server with predefined routes.
	defer server.Close()        // Ensures the server is closed once tests are done.

	// Executes a request to create an order and checks the response.
	createResp := executeRequest(server, "GET", "/create", nil)
	// Asserts that the HTTP status code for the creation response is 201 (Created).
	assert.Equal(t, http.StatusCreated, createResp.Code, "Expected status code 201 for order creation")

	// Defines a struct to unmarshal the create order response.
	var createResponse struct {
		OrderID string `json:"orderID"` // The order ID returned by the create endpoint.
	}
	// Unmarshals the JSON response into the struct and checks for errors.
	err := json.Unmarshal(createResp.Body.Bytes(), &createResponse)
	assert.NoError(t, err, "Error decoding JSON response for order creation")

	// Prints the generated order ID to the console.
	fmt.Printf("Order generated with number = %s\n", createResponse.OrderID)

	// Executes a request to retrieve the created order using its ID and checks the response.
	orderResp := executeRequest(server, "GET", "/order/"+createResponse.OrderID, nil)
	// Asserts that the HTTP status code for retrieving the order details is 200 (OK).
	assert.Equal(t, http.StatusOK, orderResp.Code, "Expected status code 200 for fetching order details")

	// Prints the details of the retrieved order to the console.
	fmt.Printf("Order details: %s\n", orderResp.Body.String())
}

// setupTestServer initializes and returns a new instance of a test HTTP server.
func setupTestServer() *httptest.Server {
	mux := http.NewServeMux() // Creates a new HTTP request multiplexer.
	// Handles the "/create" route for order creation.
	mux.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)       // Sets the response status code to 201 (Created).
		w.Write([]byte(`{"orderID": "12345"}`)) // Writes a fixed order ID as the response.
	})
	// Handles the "/order/" route for fetching order details.
	mux.HandleFunc("/order/", func(w http.ResponseWriter, r *http.Request) {
		orderID := r.URL.Path[len("/order/"):]            // Extracts the order ID from the URL.
		w.WriteHeader(http.StatusOK)                      // Sets the response status code to 200 (OK).
		w.Write([]byte(`{"orderID": "` + orderID + `"}`)) // Returns the order ID in the response.
	})
	return httptest.NewServer(mux) // Returns the test server instance.
}

// executeRequest helps in executing a request against the test server and capturing the response.
func executeRequest(server *httptest.Server, method, path string, payload []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, server.URL+path, nil) // Creates a new HTTP request with the given method, path, and payload.
	res := httptest.NewRecorder()                           // Initializes a new response recorder to capture the response.
	server.Config.Handler.ServeHTTP(res, req)               // Dispatches the request to the server's handler and records the response.
	return res                                              // Returns the recorded response.
}
