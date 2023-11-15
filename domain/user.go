package domain

type User struct {
	ID           uint
	Name         string
	PhoneNumber  string
	HashPassword string
}

type UserInfo struct {
	ID          uint
	Name        string
	PhoneNumber string
}

type UserRepo interface {
	Register(u User) (User, error)
	// Login(u User) (User, error)
	// Profile(userID uint) (User, error)
}
