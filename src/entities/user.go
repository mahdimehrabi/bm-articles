package entities

// user is only recievable from user microservice
type User struct {
	id       int64
	email    string
	fullname string
}
