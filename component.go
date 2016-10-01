package ecs

// Component is an interface which abstract each component
type Component interface {
	// ComponentType
	ComponentType() uint32
}

// Components is an interface which deals with add, remove and get componets from internal array
type Components interface {
	Add(Component)
	Get(uint32) Component
	Remove(uint32)
}

type components struct {
	components []Component
}

func (c *components) Add(component Component) {
	idx := bitCounter(component.ComponentType())
	if c.components[idx] == nil {
		c.components[idx] = component
	}
}

func (c *components) Get(uint32) Component {
	return nil
}

func (c *components) Remove(uint32) {

}

func NewComponentContainer() Components {
	return &components{
		components: make([]Component, 32),
	}
}
