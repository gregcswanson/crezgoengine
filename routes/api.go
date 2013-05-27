package routes

import (
  "net/http"
  "controllers/users"
)

func ConfigureAPI() {
  http.HandleFunc("/api/user", UserApiRoutes)
}

func UserApiRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"  {
		usersController.Current(w, r)
	} else if r.Method == "POST" {
		usersController.Current(w, r)
  	}
}