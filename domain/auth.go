package domain

type Auth struct {
	AccessToken  string
	RefreshToken string
}

type AuthGenerator interface {
	CreateAccessToken(u User) (string, error)
	CreateRefreshToken(u User) (string, error)
}
