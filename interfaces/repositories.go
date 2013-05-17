package interfaces

import (
	"domain"
  "net/http"
)

type BaseRepository struct {
	request  *http.Request
}

type UserRepositiory BaseRepository
type CommentRepositiory BaseRepository

func NewUserRepositiory(request *http.Request) *UserRepositiory {
	userRepository := new(UserRepositiory)
	userRepository.request = request
	return userRepository
}

func NewCommentRepositiory(request *http.Request) *CommentRepositiory {
	commentRepositiory := new(CommentRepositiory)
	commentRepositiory.request = request
	return commentRepositiory
}

func (repository *UserRepositiory) FindCurrent() domain.User {
  var user domain.User
  user = domain.User{}
  user.Id = 1
  user.Name = "Bob Smith"
  user.Email = "bob@smith.com"
  return user
}

func (repository *CommentRepositiory) Store(comment domain.Comment) {
	// save the comment using the storage context
}

func (repository *CommentRepositiory) FindForUserId(userid int) []domain.Comment {
  var comments []domain.Comment
	comments = make([]domain.Comment, 5)
	for i, _ := range comments {
    comments[i] = domain.Comment{i, userid, "Comment"}
	}
	return comments
}