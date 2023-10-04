package entities

type Comment struct {
	ID       int64
	UserID   int64
	Fullname string
	Body     string
}
