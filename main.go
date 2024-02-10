package main

import "fmt"

//structures in go
type sender struct {
	rateLimit int
	user
}

type user struct {
	name string 
	number int
}


//function in go
func test(s sender){
	fmt.Println("Sender name:", s.name)
	fmt.Println("Sender number:", s.number)
	fmt.Println("Sender rateLimit:", s.rateLimit)
	fmt.Println("=========================")
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error){

	costToCustomer, err := sendSMS(msgToCustomer)
     //Handling errors in go
	if err != nil {
		return 0.0, err
	}

	costForSpouse, err := sendSMS(msgToSpouse)
   //handling errors in Go
	if err != nil{
		return 0.0, err
	}

	return costToCustomer + costForSpouse , nil
}

//FIZZBUZZ
func fizzBuzz(){
   for i := 1; i<= 100; i++{
	if i%3 == 0 && i%5 == 0{
		fmt.Println("Fizzbuzz")
	}else if i%3 == 0{
       fmt.Println("fizz")
	}else if i%5 == 0{
		fmt.Println("buzz")
	}else{
		fmt.Println(i)
	}
   }
}


//Print Primes. It should print all of the prime numbers up to and including max.
//It should skip any numbers that are not prime


func printPrimes(max int){
     for n :=2; n < max+1; n++{
		if n == 2{
			fmt.Println(n)
			continue
		}
		if n%2 == 0{
			continue
		}
		isPrime := true
		for i := 3;i * i < n+1; i++{
			if n % i == 0{	
				isPrime = false
				break
			}
		}	
		if !isPrime{
			continue
		}
		fmt.Println(n)
	 }
}



func prime(max int){
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("================================")
}


func sendSMS( message string) (float64, error){
	const maxTextLen = 25
	const costPerChar = 0.0002

	if len(message)  > maxTextLen {
		return 0.0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * float64(len(message)), nil
}



func send(msgToCustomer, msgToSpouse string){
	defer fmt.Println("============")
	fmt.Println("Message for customer:", msgToCustomer)
	fmt.Println("Message for spouse:", msgToSpouse)

	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
	//handling errors in go
	if err != nil{
		fmt.Println("Error:", err)
		return 
	}

	fmt.Printf("Total cost: $%.4f", totalCost)
}


//The point of entry for execution in Go
func main() {

	x := 5

	x = increment(x)

	fmt.Println(x)

	firstName, _ := getNames()

	fmt.Println("Welcome to Textio ,", firstName)



	test(sender{
		rateLimit: 10000,
		user:user{
			name: "Deborah",
			number:18055558790,
		},
	})
	test(sender{rateLimit: 5000,
	user: user{
		name: "Sarah",
		number: 19055558790,
	},
})

test(sender{rateLimit: 1000,
	user: user{
		name: "Sally",
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

}

func increment(x int) int {	
	x++
	return x
}

func getNames() (string, string) {
	return "John", "Doe"
}
