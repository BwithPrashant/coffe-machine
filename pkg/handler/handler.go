package handler

import (
	"coffee-machine/pkg/handler/pkg/objects"
	"coffee-machine/pkg/handler/pkg/statemanager"
	"coffee-machine/pkg/handler/pkg/utils"
	"fmt"
)

const (
	// thresold percentage of an ingredient after which warning to be genarated
	THRESOLD_INGREDIENT_PERCENTAGE int = 10
)

// Init() method to be called , when coffee-machine starts
// It takes input file path and create a coffeeMachineIdentifier ,which stores metadata of coffeeMachine
// metadata is no_of_outlets, beverages and its ingredients, available ingredients and its weight etc
// For each outlet , a seperate instance of stateMachine is created. States of eachOutlet is {available | inProgress}

func Init(filepath string) error {

	var coffeMachineCreateRequest objects.CoffeMachineCreateRequest

	// Reading input file and marshal data into coffeMachineCreateRequest
	err := utils.ReadFile(filepath, &coffeMachineCreateRequest)

	// Assign thresold value of each ingredient.At any point of time, If ingredient is less than
	// its thresold value , a warning will be generated .
	// e.g. "Ingredient hot_water is running low. Actual amount: 10, Thresold amount : 50"

	AssignThresoldIngrediants(&coffeMachineCreateRequest.CoffeMachineIdentity)

	// call setter function of CoffeMachineIdentity
	objects.SetCoffeMachineIdentity(&coffeMachineCreateRequest.CoffeMachineIdentity)

	// Create statemachine for each outlet
	stateManagerMap := make(map[int]statemanager.State)
	for i := 0; i < coffeMachineCreateRequest.OutletNum; i++ {
		stateManagerMap[i] = statemanager.NewCoffeeMachineOutlet(i)
	}
	statemanager.SetStateManagerMap(stateManagerMap)
	return err
}

// Assign thresold value of each ingredient.At any point of time, If ingredient is less than
// its thresold value , a warning will be generated .
// sample warning: "Ingredient hot_water is running low. Actual amount: 10, Thresold amount : 50"
//  Thresold value is calculated as (THRESOLD_INGREDIENT_PERCENTAGE) % of total amount

func AssignThresoldIngrediants(coffeMachineIdentity *objects.CoffeMachineIdentity) {

	coffeMachineIdentity.ThresoldIngrediants = make(map[string]int)
	for ingredient, amount := range coffeMachineIdentity.Ingrediants {
		thresoldAmount := (amount * THRESOLD_INGREDIENT_PERCENTAGE) / 100
		coffeMachineIdentity.ThresoldIngrediants[ingredient] = thresoldAmount
	}
}

// Serve Beverage from a given outlet_id.
// outlet_id is sequesnce no of outlets.
// e.g. if there are 3 outlets, outlet_id is [0,1,2]

func RequestBeverages(outlet_id int, beverage_type string) error {

	// Get stateMachine for outlet_id
	// Check if outlet_id is greater than total number of outlets, if true then invalid outlet_id
	outletStateMachine, ok := statemanager.GetStateManagerMap()[outlet_id]
	if !ok {
		fmt.Printf("Invalid outlet_id %d\n", outlet_id)
		return fmt.Errorf("Invalid outlet_id %d", outlet_id)
	}

	// call ServeRequest by stateMachine instance of outlet_id
	err := outletStateMachine.ServeRequest(beverage_type)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Prepared %v from outlet %d\n", beverage_type, outlet_id)
	return nil
}

// Add ingredient to coffee machine

func AddIngredient(ingredient string, amount int) error {

	// Acquire Lock for get and update CoffeMachineIdentity
	// It is required to have lock on both get and update , otherwise
	// data can be inconsistent

	lock := objects.GetCoffeeMachineIdentityMutexLock()
	lock.Lock()
	// Release lock after function scope ends
	defer lock.Unlock()

	coffeMachineIdentity := objects.GetCoffeMachineIdentity()

	// Check if ingredient exists in coffe-machine
	ingredientMap, ok := coffeMachineIdentity.Ingrediants[ingredient]
	if !ok {
		fmt.Printf("Ingredient %v not available in coffee-machine\n", ingredient)
		return fmt.Errorf("Ingredient %v not available in coffee-machine\n", ingredient)
	}

	// update ingredient quantity
	coffeMachineIdentity.Ingrediants[ingredient] = ingredientMap + amount

	// update thresold amount of ingredient
	coffeMachineIdentity.ThresoldIngrediants[ingredient] = (coffeMachineIdentity.Ingrediants[ingredient] * THRESOLD_INGREDIENT_PERCENTAGE) / 100
	return nil
}

// List all ingredients name available in coffee machine
func ListIngrediants() []string {
	list := make([]string, 0)
	for k, _ := range objects.GetCoffeMachineIdentity().Ingrediants {
		list = append(list, k)
	}
	return list
}

// List all ingredients and its amount in coffee machine
func ListIngrediantsAvailability() map[string]int {
	data := make(map[string]int)
	for k, v := range objects.GetCoffeMachineIdentity().Ingrediants {
		data[k] = v
	}
	return data
}

// get amount of a specific ingredient
func GetIngrediantAmount(ingrediant string) int {
	return objects.GetCoffeMachineIdentity().Ingrediants[ingrediant]
}

// List all Beverages name in coffee machine
func ListBevrages() []string {
	list := make([]string, 0)
	for k, _ := range objects.GetCoffeMachineIdentity().Beverages {
		list = append(list, k)
	}
	return list
}

// get number of outlets in coffee machine
func GetOutletNums() int {
	return objects.GetCoffeMachineIdentity().OutletNum
}
