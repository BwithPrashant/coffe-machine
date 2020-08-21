package statemanager

import "sync"

// CoffeeMachineOutlet holds metadata of coffee_machine_outlet
// Stores each possible state and current state
type CoffeeMachineOutlet struct {
	availableState  State
	inProgressState State
	currentState    State
	outlet_id       int
	sync.RWMutex
}

// Returns a new instance of coffee_machine_outlet

func NewCoffeeMachineOutlet(outlet_id int) *(CoffeeMachineOutlet) {

	coffeeMachineOutlet := &CoffeeMachineOutlet{
		outlet_id: outlet_id,
	}
	coffeeMachineOutlet.availableState = &AvailableState{coffeeMachineOutlet: coffeeMachineOutlet}
	coffeeMachineOutlet.inProgressState = &InProgressState{coffeeMachineOutlet: coffeeMachineOutlet}
	coffeeMachineOutlet.SetState(coffeeMachineOutlet.availableState)
	return coffeeMachineOutlet
}

// Serve request
func (cmo *CoffeeMachineOutlet) ServeRequest(beverage string) error {
	return cmo.GetState().ServeRequest(beverage)
}

// setter function for state of coffee_machine_outlet
func (cmo *CoffeeMachineOutlet) SetState(state State) {
	cmo.Lock()
	cmo.currentState = state
	cmo.Unlock()
}

// getter function for state of coffee_machine_outlet
func (cmo *CoffeeMachineOutlet) GetState() State {
	cmo.RLock()
	defer cmo.RUnlock()
	return cmo.currentState

}

// // As it is starting with lowecase , it can't be exported to other package
var stateManagerMap map[int]State

// Getter function for state of State_manager_map
func GetStateManagerMap() map[int]State {
	return stateManagerMap
}

// Setter function for state of State_manager_map
func SetStateManagerMap(stateManagerMapParam map[int]State) {
	stateManagerMap = stateManagerMapParam
}
