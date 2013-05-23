package routes

import (
  "net/http"
  "controllers/home"
  "controllers/comments"
  "controllers/users"
)

func Configure() {
  http.HandleFunc("/", MainPage)
  http.HandleFunc("/user", usersController.Current)
  http.HandleFunc("/comments", commentsController.Index)
  http.HandleFunc("/404", homeController.FourZeroFour)
}

func MainPage(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" || r.URL.Path != "/" {
		homeController.FourZeroFour(w, r)
  } else {
    homeController.Index(w, r)
  }
  return
}