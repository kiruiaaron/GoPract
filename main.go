package main

import (
	"errors"
	"fmt"
)

func increment(x int) int {
	x++
	return x
}

func getNames() (string, string) {
	return "John", "Doe"
}

// structures in go
type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

// function in go
func test(s sender) {
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("=========================")
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {

	costToCustomer, err := sendSMS(msgToCustomer)
	//Handling errors in go
	if err != nil {
		return 0.0, err
	}

	costForSpouse, err := sendSMS(msgToSpouse)
	//handling errors in Go
	if err != nil {
		return 0.0, err
	}

	return costToCustomer + costForSpouse, nil
}

// FIZZBUZZ
func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

//Print Primes. It should print all of the prime numbers up to and including max.
//It should skip any numbers that are not prime

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

func prime(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("================================")
}

func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChar = 0.0002

	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}

func send(msgToCustomer, msgToSpouse string) {
	defer fmt.Println("============")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)

	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	//handling errors in go
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Total cost: $%.4f", totalCost)
}

//Arrays in go
//Arrays are fixed ordered list

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func sendFromArray(name string, doneAt int) {
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("complete failure")
		}
	}
}

// Slices in go
const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsupported plan")
}

// Variadic function and spread operator
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

// APPEND
// The built-in append function is used to dynamically add elements to a slice:
// Syntax: func append(slice []Type, elems ...Type) []Type
type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func appendCost(costs []cost){
	fmt.Printf("Creating daily buckets for %v costs...\n", len(costs))
	costsByday := getCostsByDay(costs)
	fmt.Println("costs by day:")
	for i := 0; i < len(costsByday); i++{
		fmt.Printf("- Day %v: %.2f\n",i, costsByday[i])
	}
	fmt.Println("===== END REPORT ======")
}


//SLICE OF SLICES

func createMatrix(rows, cols int) [][]int{
	matrix := make([][]int,0)
	for i := 0; i < rows; i++{
		row := make([]int, 0)
		for j := 0; j < cols; j++{
           row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func testMatrix(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

//Range -Go provides syntactic sugar to iterate over element of a slice
func indexOfFirstWord(msg []string, badwords []string) int{
	for i, word := range msg {
		for _, badWord := range badwords{
			if word == badWord{
				return i
			}
		}
	}
	return -1
}

//Maps
//Maps are similar to javascript objects, python dictionaries, and Ruby hashes.Maps are data structure that provides key-> value mapping.
//The zero value of map is nil.

var ages = make(map[string] int)
ages["john"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 //overwrites 24
//Alternative way 
ages = map[string]int{
	"john":37,
	"mary":21,
}

type user struct{
	names string
	phoneNumbers int
}


func getUserMap(names []string, phoneNumbers []int)(map[string]user,error){
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers){
		return nil, errors.New("Invalid sizes")
	}
	for i := i< len(names); i++{
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] =user{
			name:name,
			phoneNumber :phoneNumber,

		}
	}
	return userMap,nil
}



//Defer in go
//Defer keyword is fairly unique in go.It allows a function to be executed automatically just before its enclosing funtion returns.
func logAndDelete(users ma[string]user, name string)(log string){
	defer delete(users, name)
	user, ok := users[name]
	if !ok{
		//instead of repeatedly using delete(users, name) just use defer at the top
		return logNotFound
	}
	if user.admin{
		//instead of repeatedly using delete(users, name)  just use defer at the top
		return logAdmin
	}
	return logDeleted
}


//Closures
//A closure is a function that references variables from outside its own function body.The function may  access and assign to the referenced variables.
//In the example below, the concatter()function returns a function that has reference to an enclosed doc value. Each successive call to harryPotterAggregator mutates that same doc variable.

func concatter() func(string) string{
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}



// The point of entry for execution in Go
func main() {

	x := 5

	x = increment(x)

	fmt.Println(x)

	firstName, _ := getNames()

	fmt.Println("Welcome to Textio ,", firstName)

	test(sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	})
	test(sender{rateLimit: 5000,
		user: user{
			name:   "Sarah",
			number: 19055558790,
		},
	})

	test(sender{rateLimit: 1000,
		user: user{
			name:   "Sally",
			number: 19055558790,
		},
	})

	send(
		"Thanks for coming in to our flower shop today!",
		"We hope you enjoyed your gift.",
	)
	send(
		"Thanks for joining us!",
		"Have a good day.",
	)
	send(
		"Thank you.",
		"Enjoy!",
	)
	send(
		"We loved having you in!",
		"We hope the rest of your evening is absolutely fantastic.",
	)

	//fizzbuzz
	fizzBuzz()

	//PrintPrime
	prime(10)
	prime(20)
	prime(30)

	sendFromArray("Bob", 3)
	sendFromArray("Alice", 1)
	sendFromArray("Mangalam", 2)
	sendFromArray("Ozgur", 3)

	names := []string{"bob", "sue", "alice"}
	printStrings(names...) //spread operator


	appendCost([]cost{
		{0, 1.0},
		{1,2.0},
		{1,3.1},
		{2,2.5},
	})

	testMatrix(3,3)
	testMatrix(5,5)
	testMatrix(10,10)
	testMatrix(15,15)


	//from closures
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("drive"))
	//Mr and Mrs. Dursley of number four, Privet Drive

}
