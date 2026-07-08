package practice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type apiUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userAPI struct {
	users  []apiUserResponse
	nextID int
}

func newUserAPI() *userAPI {
	return &userAPI{
		users:  make([]apiUserResponse, 0),
		nextID: 1,
	}
}

func apiHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func (api *userAPI) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	user := apiUserResponse{
		ID:    api.nextID,
		Name:  req.Name,
		Email: req.Email,
	}
	api.nextID++
	api.users = append(api.users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (api *userAPI) listUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api.users)
}

func HTTPServersAndJSONRoutesEx16() {
	fmt.Println("1. Write a JSON response from a handler")
	fmt.Println("The handler receives a request and writes the response through w.")
	fmt.Println("---")
	fmt.Println("func apiHealthHandler(w http.ResponseWriter, r *http.Request) {")
	fmt.Println(`    w.Header().Set("Content-Type", "application/json")`)
	fmt.Println("    json.NewEncoder(w).Encode(map[string]string{")
	fmt.Println(`        "status": "ok",`)
	fmt.Println("    })")
	fmt.Println("}")
	fmt.Println("req := httptest.NewRequest(http.MethodGet, \"/health\", nil)")
	fmt.Println("rec := httptest.NewRecorder()")
	fmt.Println("apiHealthHandler(rec, req)")
	healthReq := httptest.NewRequest(http.MethodGet, "/health", nil)
	healthRec := httptest.NewRecorder()
	apiHealthHandler(healthRec, healthReq)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> status code = %d\n", healthRec.Code)
	fmt.Printf(">> content type = %s\n", healthRec.Header().Get("Content-Type"))
	fmt.Printf(">> body = %s", healthRec.Body.String())

	fmt.Println("\n2. Route requests with a mux")
	fmt.Println("A ServeMux connects method and path patterns to handlers.")
	fmt.Println("---")
	fmt.Println("api := newUserAPI()")
	fmt.Println("mux := http.NewServeMux()")
	fmt.Println(`mux.HandleFunc("GET /health", apiHealthHandler)`)
	fmt.Println(`mux.HandleFunc("GET /users", api.listUsersHandler)`)
	fmt.Println(`mux.HandleFunc("POST /users", api.createUserHandler)`)
	api := newUserAPI()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", apiHealthHandler)
	mux.HandleFunc("GET /users", api.listUsersHandler)
	mux.HandleFunc("POST /users", api.createUserHandler)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Println(">> GET /health routes to apiHealthHandler")
	fmt.Println(">> GET /users routes to api.listUsersHandler")
	fmt.Println(">> POST /users routes to api.createUserHandler")

	fmt.Println("\n3. Decode JSON and create a user")
	fmt.Println("Decode reads the request body into a struct, then Encode writes JSON back.")
	fmt.Println("---")
	fmt.Println(`body := strings.NewReader(` + "`" + `{"name":"Reda","email":"reda@example.com"}` + "`" + `)`)
	fmt.Println(`req := httptest.NewRequest(http.MethodPost, "/users", body)`)
	fmt.Println("rec := httptest.NewRecorder()")
	fmt.Println("mux.ServeHTTP(rec, req)")
	body := strings.NewReader(`{"name":"Reda","email":"reda@example.com"}`)
	createReq := httptest.NewRequest(http.MethodPost, "/users", body)
	createRec := httptest.NewRecorder()
	mux.ServeHTTP(createRec, createReq)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> status code = %d\n", createRec.Code)
	fmt.Printf(">> body = %s", createRec.Body.String())

	fmt.Println("\n4. List the users stored in memory")
	fmt.Println("The list handler encodes the current slice as JSON.")
	fmt.Println("---")
	fmt.Println(`req := httptest.NewRequest(http.MethodGet, "/users", nil)`)
	fmt.Println("rec := httptest.NewRecorder()")
	fmt.Println("mux.ServeHTTP(rec, req)")
	listReq := httptest.NewRequest(http.MethodGet, "/users", nil)
	listRec := httptest.NewRecorder()
	mux.ServeHTTP(listRec, listReq)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> status code = %d\n", listRec.Code)
	fmt.Printf(">> body = %s", listRec.Body.String())

	fmt.Println("\n5. Return a clear error for invalid input")
	fmt.Println("A handler can stop early with http.Error and a useful status code.")
	fmt.Println("---")
	fmt.Println(`body := strings.NewReader(` + "`" + `{"email":"missing-name@example.com"}` + "`" + `)`)
	fmt.Println(`req := httptest.NewRequest(http.MethodPost, "/users", body)`)
	fmt.Println("rec := httptest.NewRecorder()")
	fmt.Println("mux.ServeHTTP(rec, req)")
	missingName := strings.NewReader(`{"email":"missing-name@example.com"}`)
	errorReq := httptest.NewRequest(http.MethodPost, "/users", missingName)
	errorRec := httptest.NewRecorder()
	mux.ServeHTTP(errorRec, errorReq)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> status code = %d\n", errorRec.Code)
	fmt.Printf(">> body = %s", errorRec.Body.String())
}
