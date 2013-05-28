package domain

type UserRepository interface {
	FindCurrent() User
	Store(user User) User
  LoginUrl() (string, error)
  LogoutUrl() (string, error)
}

type CommentRepository interface {
	Store(item Comment)
	FindForUserId(userid int) []Comment
}

type User struct {
	Id   int
	Name string
  Email string
  Nickname string
  IsLoggedIn bool
}

type Comment struct {
	Id int
  UserId int
  Comment string
}