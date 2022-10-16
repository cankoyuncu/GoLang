package loops

import "fmt"

//Bravo. 3.deneme bildiniz : süper
//1-3: süper 4-6:güzel 6 ve üzeri fena değil

func Demo3() {
	myNumber := 25
	estimatedNumber := 0
	i := 0

	fmt.Println("Guess the number:")
	i = i + 1
	fmt.Scanln(&estimatedNumber)
	fmt.Println(estimatedNumber)

	for myNumber != estimatedNumber {
		if estimatedNumber > myNumber {
			fmt.Println("Make a smaller guess:")
			i = i + 1
			fmt.Scanln(&estimatedNumber)
		}
		if myNumber > estimatedNumber {
			fmt.Println("Make a bigger guess:")
			i = i + 1
			fmt.Scanln(&estimatedNumber)
		}
	}

	if estimatedNumber == myNumber {

		SuccessStatus := ""
		if i > 0 && i <= 3 {
			SuccessStatus = "Super"
		} else if i <= 6 {
			SuccessStatus = "Normal"
		} else {
			SuccessStatus = "Bad"
		}

		fmt.Printf("Right guess! %v.try find. Your success status:%v", i, SuccessStatus)
	}
}
