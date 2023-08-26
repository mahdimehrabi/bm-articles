package entities

// user is only recievable from user microservice
type User struct {
	ID       int64 `gorm:"primaryKey"`
	Email    string
	Fullname string
}
