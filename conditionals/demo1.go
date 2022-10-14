package conditionals

import "fmt"

func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 500

	if cekilmekIstenen > hesap {
		fmt.Println("Insufficient balance!")
	}
	if cekilmekIstenen <= hesap {
		fmt.Println("Your money is getting ready")
		hesap = hesap - cekilmekIstenen
	}

	fmt.Printf("Process Finish! Your balance is: %v", hesap)
}
