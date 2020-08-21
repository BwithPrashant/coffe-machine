package objects

import (
	"sync"
)

// Outlets store metadta of outlet of coffee machine
type Outlets struct {
	OutletNum int `json:"count_n"`
}

// CoffeMachineIdentity holds metadata of coffee machine
type CoffeMachineIdentity struct {
	Outlets             `json:"outlets"`
	Ingrediants         map[string]int            `json:"total_items_quantity"`
	ThresoldIngrediants map[string]int            `json:"thresoldIngrediants"`
	Beverages           map[string]map[string]int `json:"beverages"`
	sync.Mutex
}

// CoffeMachineCreateRequest holds metadata of coffee machine Create request
type CoffeMachineCreateRequest struct {
	CoffeMachineIdentity `json:"machine"`
}

// As it is starting with lowecase , it can't be exported to other package
var coffeMachineIdentity *CoffeMachineIdentity

// Getter for CoffeMachineIdentity
func GetCoffeMachineIdentity() *CoffeMachineIdentity {
	return coffeMachineIdentity
}

// Getter for mutex lock on CoffeMachineIdentity
func GetCoffeeMachineIdentityMutexLock() sync.Mutex {
	return coffeMachineIdentity.Mutex
}

// Setter of CoffeMachineIdentity
func SetCoffeMachineIdentity(coffeMachineIdentityParam *CoffeMachineIdentity) {
	coffeMachineIdentity = coffeMachineIdentityParam
}
