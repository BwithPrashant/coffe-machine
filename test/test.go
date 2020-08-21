package test

import (
	"coffee-machine/pkg/handler/pkg/handler"
	"fmt"
	"time"
)

//------------------------------------------------------------
//For each test case  , a desription of test is given inside test function
// Also output for the same test if given above the test function
//------------------------------------------------------------

// Run a test cases and comment rest
func RunTestCases() {
	//Test1()
	//Test2()
	//Test3()
	//Test4()
	//Test5()
	//Test6()
	Test7()
	//Test8()
}

/*
	Prepared hot_tea from outlet 0
	Prepared hot_coffee from outlet 1
	beverage black_tea can't be prepared because sugar_syrup is not sufficient

	beverage green_tea can't be prepared because green_mixture is not available

	Ingredient hot_milk is running low. Actual amount: 0, Thresold amount: 50
	beverage hot_coffee can't be prepared because hot_milk is not sufficient

	beverage hot_tea can't be prepared because hot_water is not sufficient

	Ingredient hot_milk is running low. Actual amount: 0, Thresold amount: 50
*/

func Test1() {
	// Get Some beverages from different outlets in sequence
	handler.RequestBeverages(0, "hot_tea")
	handler.RequestBeverages(1, "hot_coffee")
	handler.RequestBeverages(2, "black_tea")
	handler.RequestBeverages(6, "green_tea")
	handler.RequestBeverages(8, "hot_coffee")
	handler.RequestBeverages(9, "hot_tea")
}

/*
	Prepared hot_coffee from outlet 1
	Prepared hot_tea from outlet 0
	black_tea can't be prepared because sugar_syrup is not sufficient

	hot_coffee can't be prepared because sugar_syrup is not sufficient

	green_tea can't be prepared because sugar_syrup is not sufficient

	Ingredient hot_milk is running low. Actual amount: 0, Thresold amount: 50
	Ingredient hot_milk is running low. Actual amount: 0, Thresold amount: 50
*/

func Test2() {
	// get some beverages from different outlets in parallel
	go func() {
		handler.RequestBeverages(0, "hot_tea")
	}()

	go func() {
		handler.RequestBeverages(1, "hot_coffee")
	}()

	go func() {
		handler.RequestBeverages(2, "black_tea")
	}()

	go func() {
		handler.RequestBeverages(6, "green_tea")
	}()

	go func() {
		handler.RequestBeverages(8, "hot_coffee")
	}()
	time.Sleep(5 * time.Second)
}

/*
	Prepared hot_tea from outlet 1
	Prepared hot_coffee from outlet 1
	black_tea can't be prepared because hot_water is not sufficient

	green_tea can't be prepared because sugar_syrup is not sufficient

	Ingredient hot_milk is running low. Actual amount: 0, Thresold amount: 50
	hot_coffee can't be prepared because hot_milk is not sufficient

	hot_tea can't be prepared because hot_milk is not sufficient

	Ingredient hot_milk is running low. Actual amount: 0, Thresold amount: 50
*/
func Test3() {
	// Get beverages from same outlet in sequence
	handler.RequestBeverages(1, "hot_tea")
	handler.RequestBeverages(1, "hot_coffee")
	handler.RequestBeverages(1, "black_tea")
	handler.RequestBeverages(1, "green_tea")
	handler.RequestBeverages(1, "hot_coffee")
	handler.RequestBeverages(1, "hot_tea")
}

/*
	Prepared hot_tea from outlet 1
	Prepared hot_coffee from outlet 1
	black_tea can't be prepared because sugar_syrup is not sufficient

	green_tea can't be prepared because sugar_syrup is not sufficient

	Ingredient hot_milk is running low. Actual amount: 0, Thresold amount: 50
	hot_coffee can't be prepared because hot_milk is not sufficient

	hot_tea can't be prepared because hot_milk is not sufficient

	Ingredient hot_milk is running low. Actual amount: 0, Thresold amount: 50
*/

func Test4() {
	// Get beverages from same outlet in parallel
	go func() {
		handler.RequestBeverages(2, "hot_tea")
	}()

	go func() {
		handler.RequestBeverages(2, "hot_coffee")
	}()

	go func() {
		handler.RequestBeverages(2, "black_tea")
	}()

	go func() {
		handler.RequestBeverages(2, "green_tea")
	}()

	go func() {
		handler.RequestBeverages(2, "hot_coffee")
	}()
	time.Sleep(5 * time.Second)
}

/*
	Invalid outlet_id 11
	Prepared hot_coffee from outlet 2
*/
func Test5() {
	// get beverages from invalid store.
	// We have 10 outlets , try to get beverages with 11th outlet
	handler.RequestBeverages(11, "hot_tea")
	handler.RequestBeverages(2, "hot_coffee")
}

/*
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Current outlet is busy 2
	Prepared black_tea from outlet 2
*/
func Test6() {
	// Get beverages from same outlet in parallel with higher number of requests
	for i := 0; i < 20; i++ {
		go func() {
			handler.RequestBeverages(2, "black_tea")
		}()
	}
	time.Sleep(5 * time.Second)
}

/*
	Prepared latte from outlet 0
	Prepared latte from outlet 1
	latte can't be prepared because hot_water is not sufficient

	latte can't be prepared because hot_water is not sufficient

	latte can't be prepared because hot_water is not sufficient

	latte can't be prepared because hot_water is not sufficient

	latte can't be prepared because hot_water is not sufficient

	Ingredient hot_water is running low. Actual amount: 40, Thresold amount: 50
	latte can't be prepared because hot_water is not sufficient

	Prepared latte from outlet 8
	latte can't be prepared because hot_water is not sufficient

	latte can't be prepared because hot_water is not sufficient

	latte can't be prepared because hot_water is not sufficient

	Ingredient hot_water is running low. Actual amount: 10, Thresold amount: 24
*/
func Test7() {
	//Get beverages in parallel to random outlets.
	// In between add some ingredients
	// and again request more beverages

	for i := 0; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		go func(i int) {
			handler.RequestBeverages(i, "latte")
		}(i)
	}

	handler.AddIngredient("hot_water", 200)
	handler.RequestBeverages(0, "latte")
	handler.RequestBeverages(0, "latte")
}

/*
	Number of outltes are :  10
	Beverage list :  [hot_tea hot_coffee black_tea green_tea latte latte1]
	Ingredient list :  [hot_water hot_milk ginger_syrup sugar_syrup tea_leaves_syrup]
	Ingredient list with amount available:  map[ginger_syrup:100 hot_milk:500 hot_water:500 sugar_syrup:100 tea_leaves_syrup:100]
*/
func Test8() {
	fmt.Println("Number of outltes are : ", handler.GetOutletNums())
	fmt.Println("Beverage list : ", handler.ListBevrages())
	fmt.Println("Ingredient list : ", handler.ListIngrediants())
	fmt.Println("Ingredient list with amount available: ", handler.ListIngrediantsAvailability())
}
