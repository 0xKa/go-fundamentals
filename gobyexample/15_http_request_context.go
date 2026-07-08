package gobyexample

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type httpRequestContextUser struct {
	ID   int
	Name string
}

type httpRequestContextCreateUserRequest struct {
	Name string
}

var (
	httpRequestContextUsers  []httpRequestContextUser
	httpRequestContextNextID = 1
	httpRequestContextMu     sync.Mutex
)

func ShowHTTPRequestContext() {
	http.HandleFunc("/health", httpRequestContextHealthHandler)
	http.HandleFunc("/users", httpRequestContextCreateUserHandler)

	fmt.Println("server running on :8080")
	http.ListenAndServe("localhost:8080", nil)
}

func httpRequestContextHealthHandler(w http.ResponseWriter, r *http.Request) {

	r.Context()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Health check passed"))
}

func httpRequestContextCreateUserHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	// Uncomment the following lines to stimulate a timeout, we will cancel the context after 1 second
	// ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	// defer cancel()

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	user, err := httpRequestContextCreateUser(ctx, httpRequestContextCreateUserRequest{Name: name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintf(w, "created user: %+v\n", user)
}

func httpRequestContextCreateUser(ctx context.Context, req httpRequestContextCreateUserRequest) (httpRequestContextUser, error) {
	user, err := httpRequestContextInsertUser(ctx, req.Name)
	if err != nil {
		return httpRequestContextUser{}, err
	}

	return user, nil
}

func httpRequestContextInsertUser(ctx context.Context, name string) (httpRequestContextUser, error) {
	select {
	case <-time.After(2 * time.Second):
		httpRequestContextMu.Lock()
		defer httpRequestContextMu.Unlock()

		user := httpRequestContextUser{
			ID:   httpRequestContextNextID,
			Name: name,
		}

		httpRequestContextNextID++
		httpRequestContextUsers = append(httpRequestContextUsers, user)

		return user, nil

	case <-ctx.Done():
		return httpRequestContextUser{}, ctx.Err()
	}
}
