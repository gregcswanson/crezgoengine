package homeController

import (
  "fmt"
  "net/http"
  "time"
  "infrastructure"  
)

func Index(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "controllers.Index")
  message := fmt.Sprintf("Hello, %s!","Home Index")
  fmt.Fprint(w, message)
}
