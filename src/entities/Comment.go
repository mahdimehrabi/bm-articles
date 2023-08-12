package entities

type Comment struct {
	ID       int64
	UserID   int64
	User     User
	Fullname string
	Body     string
}
