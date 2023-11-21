package domain

type WaitingMember struct {
	UserID    uint
	Timestamp int64
	Category  Category
}