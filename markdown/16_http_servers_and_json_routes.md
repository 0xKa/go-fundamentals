# HTTP Servers and JSON Routes in Go

An HTTP server receives requests, chooses a handler for each route, and writes responses back to the client.

The core handler shape is:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	// read from r
	// write to w
}
```

`r` contains what the client sent. `w` is how the server sends a response.

## Handlers write responses

A handler can set response headers, choose a status code, and write a body:

```go
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
```

This response is JSON:

```json
{"status":"ok"}
```

Headers should be set before writing the body. If you do not call `WriteHeader`, Go sends `200 OK` when the first body bytes are written.

## Requests carry method, path, headers, and body

The request value has the details of the incoming HTTP request:

```go
r.Method
r.URL.Path
r.Header.Get("Content-Type")
r.Body
r.Context()
```

For a JSON API, the body is often decoded into a request struct:

```go
type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var req createUserRequest
err := json.NewDecoder(r.Body).Decode(&req)
```

`Decode` reads JSON from the request body and fills the struct. The JSON tags connect JSON field names like `"name"` to Go fields like `Name`.

## ServeMux routes requests

A `ServeMux` is Go's standard router:

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /health", healthHandler)
mux.HandleFunc("POST /users", createUserHandler)
```

The mux decides which handler should run:

```text
GET /health  -> healthHandler
POST /users  -> createUserHandler
```

Using your own mux is clearer than registering everything on the global default mux.

## Create JSON API responses

A create route usually follows this flow:

```text
decode JSON
validate input
create or store the value
write status code
encode JSON response
```

Example:

```go
type userResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	user := userResponse{ID: 1, Name: req.Name, Email: req.Email}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
```

`http.Error` writes an error response with a status code. For successful creates, `201 Created` is more specific than `200 OK`.

## Store data behind an app struct

Small examples often use package-level variables, but a tiny app struct keeps shared state explicit:

```go
type userApp struct {
	users  []userResponse
	nextID int
}
```

Handlers can be methods on that struct:

```go
func (app *userApp) listUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(app.users)
}
```

In a real server, many requests can run at the same time. Shared state like a slice should be protected with a mutex or moved into a database-backed service.

## Test handlers without opening a port

The `net/http/httptest` package lets you call handlers directly:

```go
req := httptest.NewRequest(http.MethodGet, "/health", nil)
rec := httptest.NewRecorder()

healthHandler(rec, req)

fmt.Println(rec.Code)
fmt.Println(rec.Body.String())
```

This is useful for practice and for real tests. It checks the same handler code without starting a long-running server.

## Summary

HTTP servers in Go are built from handlers, a mux that routes requests, and disciplined request/response code.

- A handler receives `http.ResponseWriter` and `*http.Request`.
- Use `http.NewServeMux` to register routes explicitly.
- Decode JSON request bodies into structs and encode response structs back to JSON.
- Set headers and status codes before writing the response body.
- Use `httptest` to practice and test handlers without binding to a port.
