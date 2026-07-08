package gobyexample

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

// what to receive from the client
type httpServerExampleCreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// what to send back to the client
type httpServerExampleUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// struct for the app state
type httpServerExampleApp struct {
	logger *slog.Logger                    // logs
	mu     sync.Mutex                      // protects the users slice and nextID from concurrent access
	users  []httpServerExampleUserResponse // list of users (stored in-memory)
	nextID int                             // next user ID to assign
}

func ShowHTTPServerExample() {
	// create a logger that writes to stdout (console), it can be replaced with a file or other destinations
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// initialize the app state
	app := &httpServerExampleApp{
		logger: logger,
		users:  make([]httpServerExampleUserResponse, 0),
		nextID: 1,
	}

	// use mux to route requests to the appropriate handlers
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", httpServerExampleHealthHandler)
	mux.HandleFunc("GET /users", app.listUsersHandler)
	mux.HandleFunc("POST /users", app.createUserHandler)

	// create an HTTP server and configure it
	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// log that the server is starting
	logger.Info("server running", "addr", srv.Addr)

	// start the server and log any errors if it stops unexpectedly
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}

// handlers

// (/health) endpoint
func httpServerExampleHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

// (/users) endpoint to list all users
func (a *httpServerExampleApp) listUsersHandler(w http.ResponseWriter, r *http.Request) {

	// lock the mutex to safely access the users slice and copy it to a new slice to avoid race conditions, this allow sync access to the users slice while still allowing other goroutines to read from it without blocking
	a.mu.Lock()
	users := make([]httpServerExampleUserResponse, len(a.users))
	copy(users, a.users)
	a.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// (/users) endpoint to create a new user
func (a *httpServerExampleApp) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req httpServerExampleCreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	a.mu.Lock()
	user := httpServerExampleUserResponse{
		ID:    a.nextID,
		Name:  req.Name,
		Email: req.Email,
	}
	a.nextID++
	a.users = append(a.users, user)
	a.mu.Unlock()

	a.logger.Info("user created", "id", user.ID, "email", user.Email)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
