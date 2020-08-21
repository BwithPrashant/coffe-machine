package statemanager

// State holds each action to be performed by each actor(state instance)

type State interface {
	// Serve a beverage based on availability in coffee machine
	ServeRequest(beverage string) error
}
