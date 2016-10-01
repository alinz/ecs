package ecs

// Manager is an interface for every defined managers
type Manager interface {
	ManagerType() uint32
	Initialize()
	Cleanup()
}
