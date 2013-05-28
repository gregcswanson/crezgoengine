package homeController

import (
  "net/http"
  "time"
  "infrastructure" 
  "usecases"
  "interfaces"
  "infrastructure/views"
)

func CreateUserInteractor(r *http.Request) *usecases.UserInteractor {
  userInteractor := new(usecases.UserInteractor)
	userInteractor.UserRepository = interfaces.NewUserRepositiory(r)
  return userInteractor
}

func Index(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "homeController.Index")
  userInteractor := CreateUserInteractor(r)
  if userInteractor.IsLoggedIn() {
    _, v := views.Static("index")
    w.Write(v)
  } else {
    _, v := views.Static("landing")
    w.Write(v)
  }
}

func Login(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "homeController.Login")
  userInteractor := CreateUserInteractor(r)
  url, err := userInteractor.LoginUrl()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Location", url)
  w.WriteHeader(http.StatusFound)
  return
}

func Logout(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "homeController.Logout")
  userInteractor := CreateUserInteractor(r)
  url, err := userInteractor.LogoutUrl()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Location", url)
  w.WriteHeader(http.StatusFound)
  return
}
