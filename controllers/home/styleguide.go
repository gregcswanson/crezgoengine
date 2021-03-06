package homeController

import (
  "net/http"
  "time"
  "infrastructure"
  "infrastructure/views"
)

func Styleguide(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "controllers.styleguide")
    _, v := views.Static("styleguide")
  w.Write(v)
}