package domain

type User struct {
	ID           uint
	Name         string
	PhoneNumber  string
	HashPassword string
	Role         Role
}

// TODO - domain module could separated into two pkg entity and repository
type UserRepo interface {
	Register(u User) (User, error)
	GetUserByPhoneNumber(phoneNumber string) (User, error)
	GetUserByID(userID uint) (User, error)
}
