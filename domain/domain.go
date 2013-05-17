package domain

type UserRepository interface {
	FindCurrent() User
}

type CommentRepository interface {
	Store(item Comment)
	FindForUserId(userid int) []Comment
}

type User struct {
	Id   int
	Name string
  Email string
}

type Comment struct {
	Id int
  UserId int
  Comment string
}