package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"testcodex/models"
	"testcodex/services"
)

// UserController handles HTTP requests for users.
type UserController struct {
	service *services.UserService
}

// NewUserController creates a new controller.
func NewUserController(s *services.UserService) *UserController {
	return &UserController{service: s}
}

// HandleUsers handles "/users" route for GET and POST.
func (uc *UserController) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		uc.listUsers(w, r)
	case http.MethodPost:
		uc.createUser(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// HandleUserByID handles "/users/{id}" routes.
func (uc *UserController) HandleUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		uc.getUser(w, r, id)
	case http.MethodPut:
		uc.updateUser(w, r, id)
	case http.MethodDelete:
		uc.deleteUser(w, r, id)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (uc *UserController) createUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	created := uc.service.Create(&u)
	respondJSON(w, http.StatusCreated, created)
}

func (uc *UserController) listUsers(w http.ResponseWriter, r *http.Request) {
	users := uc.service.List()
	respondJSON(w, http.StatusOK, users)
}

func (uc *UserController) getUser(w http.ResponseWriter, r *http.Request, id int) {
	user, err := uc.service.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, user)
}

func (uc *UserController) updateUser(w http.ResponseWriter, r *http.Request, id int) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	updated, err := uc.service.Update(id, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, updated)
}

func (uc *UserController) deleteUser(w http.ResponseWriter, r *http.Request, id int) {
	if err := uc.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
