package inventory

import (
	"coffee-machine/pkg/handler/pkg/objects"
	"fmt"
	"time"
)

// Check and generate warning if amount of any ingredient is less than it's thresold amount
// A defer function is added in main().which writes "true" to "stopChannel" channel.
// When this GoRoutine/thread receives a messege on channel , it return from the function
// In case if program finishes abruptly,  defer function is called and below infinite loop breaks and avoid memory leaks

func CheckAvailability(stopChannel chan bool) {
	for {
		select {
		case <-stopChannel:
			break
		default:
			// audit for warning messege every 5 seconds
			coffeeMachineIdentity := objects.GetCoffeMachineIdentity()
			coffeeMachineIdentity.Lock()
			for k, v := range coffeeMachineIdentity.ThresoldIngrediants {
				if coffeeMachineIdentity.Ingrediants[k] < v {
					fmt.Printf("Ingredient %s is running low. Actual amount: %d, Thresold amount: %d\n", k, coffeeMachineIdentity.Ingrediants[k], v)
				}
			}

			coffeeMachineIdentity.Unlock()
			time.Sleep(5 * time.Second)
		}
	}
}

// 1. it takes beverage_name as input
// 2. Check if all required ingredient is available
// 3. If available update the quantity of ingredients in coffee machine, else error out

func UpdateIngredientAvailability(beverage string) error {
	// Acquire Lock for get and update CoffeMachineIdentity
	// It is required to have lock on both get and update , otherwise
	// data can be inconsistent

	lock := objects.GetCoffeeMachineIdentityMutexLock()
	lock.Lock()

	// Release lock after function scope ends
	defer lock.Unlock()

	coffeMachineIdentity := objects.GetCoffeMachineIdentity()

	// check if beverage exists in coffe machine
	beverageMap, ok := coffeMachineIdentity.Beverages[beverage]
	if !ok {
		return fmt.Errorf("beverage %s not found in coffee machine\n", beverage)
	}

	updatedIngredient := make(map[string]int)
	// loop through each ingredients required for beverage and check the availbility
	for k, v := range coffeMachineIdentity.Ingrediants {
		updatedIngredient[k] = v
	}
	//updatedIngredient := coffeMachineIdentity.Ingrediants
	for ingredient, amount := range beverageMap {
		availableAmount, ok := coffeMachineIdentity.Ingrediants[ingredient]
		if !ok {
			return fmt.Errorf("%s can't be prepared because %s is not available\n", beverage, ingredient)
		}
		if amount > availableAmount {
			return fmt.Errorf("%s can't be prepared because %s is not sufficient\n", beverage, ingredient)
		}
		updatedIngredient[ingredient] = (availableAmount - amount)
	}

	// If all required ingredients is available , update the amount of ingredient in coffee machine
	coffeMachineIdentity.Ingrediants = updatedIngredient
	objects.SetCoffeMachineIdentity(coffeMachineIdentity)
	return nil
}
