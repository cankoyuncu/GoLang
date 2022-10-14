package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000

	if cekilmekIstenen > hesap {
		fmt.Println("Insufficient balance!")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Your money is getting ready")
		fmt.Println("New account balance: $0")

	} else {
		fmt.Println("Your money is getting ready")
	}
}
