package variables

import "fmt"

//PascalCase
//camelCase

func Demo1() {

	var text string = "Hello Turkey!"
	fmt.Println(text)
	fmt.Println(text)
	fmt.Println(text)

	fmt.Print(text)
	fmt.Println(text)

	var kdv int = 10
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))
	fmt.Println()

	var kdv2 float32 = 0.5
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)

	kdv3 := 25.5
	fmt.Println(kdv3)
	fmt.Printf("veri tipi : %T\n", kdv3)

	var durum bool

	var text1 string = "Jack"
	var text2 string = "London"

	durum = text1 != text2

	fmt.Println(durum)
	fmt.Println(!durum)
}
