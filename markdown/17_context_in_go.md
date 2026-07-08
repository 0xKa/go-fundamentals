# Context in Go

In short:

> context is a value you pass down so inner services can stop work when the caller no longer needs the result.

In details:

A `context.Context` carries the lifetime of some work. For your first mental model, keep it simple:

```text
context = request lifetime + cancellation signal
```

It helps your code notice:

```text
this request is still alive
this request was canceled
this work should stop
```

Context does not stop your code by itself. It gives your code, and the libraries you call, a signal that the work is no longer needed.

## Start with `r.Context()`

In an HTTP handler, the most important context usually comes from the request:

```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// pass ctx into the next function
}
```

That context is tied to the request lifetime. If the client disconnects or the request is canceled, `ctx` is canceled too.

## Pass context down

Production Go usually passes context as the first argument:

```go
func createUser(ctx context.Context, req CreateUserRequest) (User, error) {
	// work here
}
```

The flow often looks like this:

```text
handler
-> service
-> database or API call
```

For practice, an in-memory slice can stand in for a database:

```go
type userStore struct {
	users  []User
	nextID int
}
```

The handler gets the request context, then passes it down:

```go
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := createUser(ctx, store, CreateUserRequest{
		Name: r.URL.Query().Get("name"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusRequestTimeout)
		return
	}

	fmt.Fprintf(w, "created user: ID=%d Name=%s\n", user.ID, user.Name)
}
```

The service passes the same context into the store:

```go
func createUser(ctx context.Context, store *userStore, req CreateUserRequest) (User, error) {
	return store.insertUser(ctx, req.Name)
}
```

This keeps the request lifetime connected to the deeper work.

## Stop slow work with `ctx.Done()`

`ctx.Done()` is a channel that closes when the context is canceled.

A slow insert can use `select` to wait for either the work or the cancellation signal:

```go
func (store *userStore) insertUser(ctx context.Context, name string) (User, error) {
	select {
	case <-time.After(10 * time.Millisecond):
		user := User{ID: store.nextID, Name: name}
		store.nextID++
		store.users = append(store.users, user)
		return user, nil

	case <-ctx.Done():
		return User{}, ctx.Err()
	}
}
```

If the work finishes first, the user is saved to the slice. If the request context is canceled first, the function returns `ctx.Err()` and does not save the user.

Common errors are:

```text
context.Canceled
context.DeadlineExceeded
```

For now, focus on the shape more than the exact error:

```text
r.Context()
-> pass ctx as the first argument
-> check ctx.Done() in slow work
-> return ctx.Err() when canceled
```

## Summary

Context is how Go connects the lifetime of an HTTP request to the work that request starts.

- Start request work with `ctx := r.Context()`.
- Pass `ctx` down as the first parameter.
- Use the same `ctx` in service and storage functions.
- Use `ctx.Done()` when slow work should stop if the request is canceled.
- Return `ctx.Err()` to explain why the context stopped.
