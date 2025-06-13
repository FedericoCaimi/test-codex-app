# Test Codex App

This is a simple REST API written in Go. It demonstrates a layered architecture with routes, controllers, services and clients.

## Running

```
go run ./...
```

The server listens on `:8080` and exposes the following endpoints:

- `POST /users` – create a new user
- `GET /users` – list all users
- `GET /users/{id}` – retrieve a user by ID
- `PUT /users/{id}` – update a user
- `DELETE /users/{id}` – delete a user

Data is stored in memory via an in-memory cache client.

