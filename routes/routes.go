package routes

import (
	"net/http"

	"testcodex/controllers"
)

// SetupRoutes configures the HTTP routes.
func SetupRoutes(controller *controllers.UserController) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", controller.HandleUsers)
	mux.HandleFunc("/users/", controller.HandleUserByID)
	return mux
}
