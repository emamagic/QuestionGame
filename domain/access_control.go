package domain

// AccessControl only keeps allowed permissions
type AccessControl struct {
	ID           uint
	ActorID      uint
	ActorType    ActorType
	PermissionID uint
}

type ActorType string

const (
	RoleActorType = "role"
	UserActorType = "user"
)