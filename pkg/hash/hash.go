package hash

type HashPassGen interface {
	GenerateFromPassword(password string) (string, error)
}

type HashPassCompare interface {
	CompareHashAndPassword(hashPassword string, password string) error
}