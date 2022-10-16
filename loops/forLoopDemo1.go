package loops

import "fmt"

func Demo1() {
	var text string = "Hello World!"

	i := 1

	for i <= 10 {
		fmt.Println(text)
		i = i + 1
	}

}
