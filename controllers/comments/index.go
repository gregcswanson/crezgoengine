package commentsController

import (
  "fmt"
  "net/http"
  "time"
  "usecases"
  "interfaces"
  "infrastructure"  
)

func Index(w http.ResponseWriter, r *http.Request) {
  defer infrastructure.TimeTrack(time.Now(), "CommentController.Index")
  // dependancy injection
  commentInteractor := new(usecases.CommentInteractor)
	commentInteractor.UserRepository = interfaces.NewUserRepositiory(r)
	commentInteractor.CommentRepository = interfaces.NewCommentRepositiory(r)
  // get the data
  comments, _ := commentInteractor.Comments()
  message := "User Comment:"
  for _, comment := range comments {
    message = message + comment.Comment + "; "
	}
  fmt.Fprint(w, message)
}