package practice

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

type contextUser struct {
	ID   int
	Name string
}

type contextCreateUserRequest struct {
	Name string
}

type contextUserStore struct {
	users  []contextUser
	nextID int
}

func newContextUserStore() *contextUserStore {
	return &contextUserStore{
		users:  make([]contextUser, 0),
		nextID: 1,
	}
}

func (store *contextUserStore) createUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	user, err := createContextUser(ctx, store, contextCreateUserRequest{Name: name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintf(w, "created user: ID=%d Name=%s\n", user.ID, user.Name)
}

func createContextUser(ctx context.Context, store *contextUserStore, req contextCreateUserRequest) (contextUser, error) {
	return store.insertUser(ctx, req.Name)
}

func (store *contextUserStore) insertUser(ctx context.Context, name string) (contextUser, error) {
	select {
	case <-time.After(10 * time.Millisecond):
		user := contextUser{
			ID:   store.nextID,
			Name: name,
		}
		store.nextID++
		store.users = append(store.users, user)
		return user, nil

	case <-ctx.Done():
		return contextUser{}, ctx.Err()
	}
}

func ContextInGoEx17() {
	store := newContextUserStore()

	fmt.Println("1. Start with the request context")
	fmt.Println("An HTTP handler gets the context from the incoming request.")
	fmt.Println("---")
	fmt.Println("func (store *contextUserStore) createUserHandler(w http.ResponseWriter, r *http.Request) {")
	fmt.Println("    ctx := r.Context()")
	fmt.Println(`    name := r.URL.Query().Get("name")`)
	fmt.Println("    user, err := createContextUser(ctx, store, contextCreateUserRequest{Name: name})")
	fmt.Println("    _ = user")
	fmt.Println("    _ = err")
	fmt.Println("}")
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Println(">> ctx now represents this request's lifetime")

	fmt.Println("\n2. Pass context from handler to service to store")
	fmt.Println("The same ctx travels through each layer that does the work.")
	fmt.Println("---")
	fmt.Println("func createContextUser(ctx context.Context, store *contextUserStore, req contextCreateUserRequest) (contextUser, error) {")
	fmt.Println("    return store.insertUser(ctx, req.Name)")
	fmt.Println("}")
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Println(">> handler -> service -> in-memory user store")

	fmt.Println("\n3. Create a user while the request is alive")
	fmt.Println("The store uses an in-memory slice as the database.")
	fmt.Println("---")
	fmt.Println(`req := httptest.NewRequest(http.MethodGet, "/users?name=Reda", nil)`)
	fmt.Println("rec := httptest.NewRecorder()")
	fmt.Println("store.createUserHandler(rec, req)")
	req := httptest.NewRequest(http.MethodGet, "/users?name=Reda", nil)
	rec := httptest.NewRecorder()
	store.createUserHandler(rec, req)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> status code = %d\n", rec.Code)
	fmt.Printf(">> body = %s", rec.Body.String())

	fmt.Println("\n4. Stop the insert when the context is canceled")
	fmt.Println("A canceled context makes ctx.Done() ready before the slow insert finishes.")
	fmt.Println("---")
	fmt.Println("ctx, cancel := context.WithCancel(context.Background())")
	fmt.Println("cancel()")
	fmt.Println(`req := httptest.NewRequest(http.MethodGet, "/users?name=Ada", nil).WithContext(ctx)`)
	fmt.Println("rec := httptest.NewRecorder()")
	fmt.Println("store.createUserHandler(rec, req)")
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	canceledReq := httptest.NewRequest(http.MethodGet, "/users?name=Ada", nil).WithContext(ctx)
	canceledRec := httptest.NewRecorder()
	store.createUserHandler(canceledRec, canceledReq)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> status code = %d\n", canceledRec.Code)
	fmt.Printf(">> body = %s", canceledRec.Body.String())
	fmt.Printf(">> saved users = %d\n", len(store.users))
}
