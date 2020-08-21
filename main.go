package main

import (
	"coffee-machine/pkg/handler/pkg/handler"
	"coffee-machine/pkg/handler/pkg/inventory"
	"coffee-machine/pkg/handler/pkg/utils"
	"coffee-machine/pkg/handler/test"
	"fmt"
)

func main() {
	fmt.Println("-------------------coffee-machine-----------------------\n")

	stopChannel := make(chan bool)

	err := handler.Init("input.json")
	if err != nil {
		fmt.Printf("Error in starting coffee machine. Error: %v\n", err)
		return
	}
	go inventory.CheckAvailability(stopChannel)
	defer utils.CloseChannel(stopChannel)
	test.RunTestCases()
}
