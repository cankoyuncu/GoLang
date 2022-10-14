package conditionals

import "fmt"

func Demo3() {
	//3 adet int degisken tanımla --sayı
	//Ekrana en buyuk olanı yazdır

	var number1, number2, number3 int = 100, 2000, 30

	var big int = number1

	if number2 > big {
		big = number2
	}
	if number3 > big {
		big = number3
	}

	fmt.Printf("Big number is: %v", big)
}
