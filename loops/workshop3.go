package loops

import "fmt"

//220 and 284 arkadas sayilardir
//10 ve 65 arkadas sayi degildir

func Demo5() {
	number1 := 220
	number2 := 284

	total1 := 0
	total2 := 0

	for i := 1; i < number1; i++ {
		if number1%i == 0 {
			total1 = total1 + i
		}
	}

	for i2 := 1; i2 < number2; i2++ {
		if number2%i2 == 0 {
			total2 = total2 + i2
		}
	}

	if total1 == number2 && total2 == number1 {
		fmt.Println("Friendly number")
	} else {
		fmt.Println("Not friendly number")
	}

}
