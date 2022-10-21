package loops

import "fmt"

// User enter the number
// 23: Prime
func Demo4() {
	number := 0
	fmt.Println("Enter the number:")
	fmt.Scanln(&number)

	prime := true
	for i := 2; i < number; i++ {
		if number%i == 0 {
			prime = false
		}
	}

	if prime == true {
		fmt.Println("Prime Nubmer")
	} else {
		fmt.Println("Is not a prime number")
	}
}
