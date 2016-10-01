package ecs

// System should be stateless if possible to make things
// simpler. the Update will be called by World implementation. the delta is difference
// between two time frame between previous call and current call.
type System interface {
	Update(delta float32, world World)
}
