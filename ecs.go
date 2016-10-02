package ecs

// Component represents the abstract version of data.
// each component will associate with a meaningful data.
// because of underneath structure, we can have upto
// 32 components in total per game.
// hope that enough...
type Component interface {
	ComponentType() uint32
}

// Entity is a shell around multiple components.
// components can be added or removed from entity either at
// runtime or compile time.
type Entity interface {
	Component(typ uint32) Component
	AddComponent(component Component)
	RemoveComponent(componentType uint32)
	HasComponentTypes(componentTypes uint32) bool
}

// Query is a bridge between System and Entity.
// query can be used to fetch array of entities
// based on the data that they have or accessing
// an already exist system.
type Query interface {
	Entities(componentTypes uint32) []Entity
	AddEntity(Entity)
	RemoveEntity(Entity)
	System(systemType uint32) System
}

// System contains all the methods which will needed for
// hopfully all the scenarioes. There is an Update method
// that is being called by World and was given a Query object
// for manipulating entities.
type System interface {
	SystemType() uint32
	Active() bool
	Start()
	Pause()
	Resume()
	End()
	Update(delta float32, query Query)
}

// World is a simple abstraction of minimum requirement of
// method which needed for any game.
type World interface {
	AddSystem(system System)
	RemoveSystem(systemType uint32)
	Update(delta float32)
}
