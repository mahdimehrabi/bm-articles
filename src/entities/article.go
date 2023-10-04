package entities

type Article struct {
	ID       int64 `gorm:"primaryKey"`
	Title    string
	Body     string
	Tags     string
	Price    float64
	BuyCount int64
	UserID   int64
}
