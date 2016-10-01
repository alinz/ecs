package ecs

// Entity is the interface for holding components
type Entity interface {
	ID() uint64
	Component(uint32) Component
}
