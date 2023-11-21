package domain

type Category string

const (
	FootballCategory = "football"
)

func (c Category) IsValid() bool {
	switch c {
	case FootballCategory:
		return true
		
	}
	return false
}
