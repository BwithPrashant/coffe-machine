COFFE-MACHINE 

Prerequisites 
1. Golang should be installed

Highlights:
1. We are taking input from input.json file and storing machine information in struct
2. We exposed all required functions in "handler" package , which is used by end customer
    functions are:
    func RequestBeverages(outlet_id int, beverage_type string) error
    func AddIngredient(ingredient string, amount int) error
    func ListIngrediants() []string
    func ListIngrediantsAvailability() map[string]int
    func GetIngrediantAmount(ingrediant string) int
    func ListBevrages() []string
    func GetOutletNums() int

3. Apart from problem statement , i am assuming , client will provide outlet_id and beverage type.
    where order_id is sequence number of outlet(0,1,2,3...n-1) , n= number of outlets

4. There can be parallel request from machine. 
    So we have maintained state of each outlet and assigned instace of statemachine for each outlet.
    At any point of time, an outlet can be in two state 
    a. Avaialbale
    b. In Progress

    As an enhancement , we can add more states in order to display more releavent information

5. When a beverage is requested from a outlet, corresponding statemachine comes into execution.
    If it is in available state , it grant the request and check for availability of all required ingredient in inventory.
        If all ingredients are available , inventory is updated and response messege is sent.
        If ingredients are not available , Then an error messesge is sent
    If outlet is in progress state , then an error messege is sent.

6. For data consistency and concurrency , we have used mutex lock.
    It can be further enhanced , when actual database comes into the picture.
    We can utilize concurrency of database. e.g. Some database offers locking mechanism over a key stored.

7. For auditing purpose, we have created a goroutine(thread) , which continuosly looks into inventory every 5 seconds.
    When machine starts , we assign some thresold amount for each ingredient. Currently it is 10 % of total amount.
    When the audit thread finds that amount of an ingredient is less than the thresold value, it generates a warnign and print
    warning to console.

8. For refilling ingredients , there is a separate function , which takes ihgredient name and amount. and add to the inventory.

---------------------------------
-----------How to Run------------
---------------------------------

copy the project folder to any directory
project folder name : coffee-machine

cd coffee-machine
go run main.go 

For running different testcases

All test cases are available in coffee-machine/test/test.go
Please comment or uncomments the testcase your are interested in



