package routes

import (
  "net/http"
  "controllers/home"
  "controllers/comments"
  "controllers/users"
)

func Configure() {
    http.HandleFunc("/", homeController.Index)
    http.HandleFunc("/user", usersController.Current)
    http.HandleFunc("/comments", commentsController.Index)
}