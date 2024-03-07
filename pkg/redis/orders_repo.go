// Package redis provides functionality for interacting with a Redis database.
package redis

import (
	"context"          // Used for managing deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	"encoding/json"    // Provides functionality for encoding and decoding JSON data.
	"store/pkg/domain" // Imports the domain package where the Order struct is defined.

	"github.com/go-redis/redis/v9" // Redis client for Go.
	"github.com/sirupsen/logrus"   // A logging library.
)

// OrdersRepo struct holds a client to the Redis database and a logger.
type OrdersRepo struct {
	c      *redis.Client  // Redis client for connecting to the Redis server.
	logger *logrus.Logger // Logger for logging error or info messages.
}

// NewOrdersRepo creates and returns a new OrdersRepo instance.
// It accepts the Redis server address and a logger as arguments.
func NewOrdersRepo(addr string, logger *logrus.Logger) *OrdersRepo {
	c := redis.NewClient(&redis.Options{Addr: addr}) // Initializes a new Redis client.
	return &OrdersRepo{c: c, logger: logger}         // Returns a new OrdersRepo with the initialized Redis client and logger.
}

// Save serializes the order object into JSON and saves it to Redis using the order ID as the key.
// It accepts a context and an order object. Returns an error if the operation fails.
func (r *OrdersRepo) Save(ctx context.Context, order domain.Order) error {
	serializedOrder, err := json.Marshal(order) // Serialize the order object into JSON.
	if err != nil {
		// Log the error if serialization fails.
		r.logger.WithFields(logrus.Fields{"orderID": order.ID, "error": err.Error()}).Error("Error serializing order")
		return err // Return the error.
	}

	// Save the serialized order to Redis with the order ID as the key.
	err = r.c.Set(ctx, order.ID, serializedOrder, 0).Err()
	if err != nil {
		// Log the error if saving to Redis fails.
		r.logger.WithFields(logrus.Fields{"orderID": order.ID, "error": err.Error()}).Error("Error saving order to Redis")
	}
	return err // Return the error (nil if operation was successful).
}

// Get retrieves an order from Redis by its ID and deserializes it into an Order object.
// It accepts a context and an order ID. Returns the Order object and an error if the operation fails.
func (r *OrdersRepo) Get(ctx context.Context, id string) (domain.Order, error) {
	serializedOrder, err := r.c.Get(ctx, id).Result() // Retrieve the order from Redis by ID.
	if err != nil {
		// Log the error if retrieval fails.
		r.logger.WithFields(logrus.Fields{"orderID": id, "error": err.Error()}).Error("Error retrieving order from Redis")
		return domain.Order{}, err // Return an empty Order object and the error.
	}
	var order domain.Order
	err = json.Unmarshal([]byte(serializedOrder), &order) // Deserialize the JSON into an Order object.
	if err != nil {
		// Log the error if deserialization fails.
		r.logger.WithFields(logrus.Fields{"orderID": id, "error": err.Error()}).Error("Error deserializing order")
		return domain.Order{}, err // Return an empty Order object and the error.
	}
	return order, nil // Return the deserialized Order object and nil (no error).
}
