package ecs

type world struct {
	systemTypes uint32
	systems     []System
	entities    []Entity
}

func (w *world) indexOfSystem(systemType uint32) int {
	position := calcBitIndex(w.systemTypes, systemType)
	return int(position - 1)
}

func (w *world) AddSystem(system System) {
	//make sure that system always passes as non nil value
	if system == nil {
		return
	}

	systemType := system.SystemType()

	//system already inside this entity32
	if w.systemTypes&systemType != 0 {
		return
	}

	w.systemTypes = w.systemTypes | systemType

	index := w.indexOfSystem(systemType)

	//insert the new system into right index
	w.systems = w.systems[0 : len(w.systems)+1]
	copy(w.systems[index:], w.systems[index+1:])
	w.systems[index] = system
}

func (w *world) RemoveSystem(systemType uint32) {
	//system doesn't have that system
	if w.systemTypes&systemType == 0 {
		return
	}

	index := w.indexOfSystem(systemType)

	//deleting the system from list
	copy(w.systems[index:], w.systems[index+1:])
	w.systems[len(w.systems)-1] = nil
	w.systems = w.systems[:len(w.systems)-1]
}

func (w *world) Update(delta float32) {
	for _, system := range w.systems {
		if system.Active() {
			system.Update(delta, w)
		}
	}
}

func (w *world) Entities(componentTypes uint32) []Entity {
	var entities []Entity

	for _, entity := range w.entities {
		if entity.HasComponentTypes(componentTypes) {
			entities = append(entities, entity)
		}
	}

	return entities
}

func (w *world) System(systemType uint32) System {
	return nil
}

func (w *world) AddEntity(entity Entity) {
	w.entities = append(w.entities, entity)
}

func (w *world) RemoveEntity(target Entity) {
	index := -1
	for i, entity := range w.entities {
		if entity == target {
			index = i
			break
		}
	}

	if index == -1 {
		return
	}

	copy(w.entities[index:], w.entities[index+1:])
	w.entities[len(w.entities)-1] = nil
	w.entities = w.entities[:len(w.entities)-1]
}

func NewWorld() World {
	return &world{}
}
