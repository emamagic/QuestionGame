package hash

type Service interface {
	GenerateFromPassword(password string) (string, error)
	CompareHashAndPassword(hashPassword string, password string) error
}
