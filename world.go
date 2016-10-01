package ecs

// World this is world object
type World interface {
	AddManager(Manager)
	Entities(Query) []Entity
	Manager(Query) Manager
	Initialize()
	Pause()
	Resume()
	Cleanup()
}

type world struct {
}

func (w *world) AddManager(Manager) {

}

func (w *world) Entities(Query) []Entity {
	return nil
}

func (w *world) Manager(Query) Manager {
	return nil
}

func (w *world) Initialize() {

}

func (w *world) Pause() {

}

func (w *world) Resume() {

}

func (w *world) Cleanup() {

}
