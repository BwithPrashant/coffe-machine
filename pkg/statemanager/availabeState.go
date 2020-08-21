package statemanager

import (
	"coffee-machine/pkg/handler/pkg/inventory"
	"time"
)

// State instance of cooffe_machine_outlet
type AvailableState struct {
	coffeeMachineOutlet *CoffeeMachineOutlet
}

// Serve beverage if coffee_machine_outlet is in available state
func (as *AvailableState) ServeRequest(beverage string) error {

	// Mark coffee_machine_outlet as inProgress state
	as.coffeeMachineOutlet.SetState(as.coffeeMachineOutlet.inProgressState)

	// check and update  amount of ingredients required to serve beverage
	// Adding a time of 1 second , so that more feasible test cases can be added

	time.Sleep(time.Second)
	err := inventory.UpdateIngredientAvailability(beverage)

	// After completing request , mark coffee_machine_outlet in avaialable state
	as.coffeeMachineOutlet.SetState(as.coffeeMachineOutlet.availableState)
	if err != nil {
		return err
	}
	return nil
}
