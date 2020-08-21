package statemanager

import (
	"fmt"
	"time"
)

// State instance of cooffe_machine_outlet
type InProgressState struct {
	coffeeMachineOutlet *CoffeeMachineOutlet
}

// Serve beverage if coffee_machine_outlet is in inProgress state
func (ips *InProgressState) ServeRequest(beverage string) error {

	// Adding time of 1 milisecond , so that more feasible test cases can be added
	time.Sleep(1 * time.Millisecond)
	// No request can be served in inProgress state
	return fmt.Errorf("Current outlet is busy %d", ips.coffeeMachineOutlet.outlet_id)
}
