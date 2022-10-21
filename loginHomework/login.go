package login

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Demo1() {

	var username string = "deneme"
	var password string = "deneme123"

	var username2 string = ""
	var password2 string = ""

	status := false

	fmt.Println("Welcome! Please enter the username and password.")
	fmt.Println("(You have only 5 right to try!)")

	for i := 4; i >= 0; i-- {
		fmt.Print("Enter the username:")
		fmt.Scanln(&username2)

		fmt.Print("Enter the password:")
		fmt.Scanln(&password2)

		if username == username2 && password == password2 {

			fmt.Println("Login successful...")

			status = true

		} else {

			fmt.Printf("Login failed...Your remaining right: %v \n", i)
		}

		file, err := os.OpenFile("./login/log.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 755)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		time1 := time.Now()
		time2 := fmt.Sprintf("%d-%d-%d	%d:%d:%d ", time1.Year(),
			time1.Month(),
			time1.Day(),
			time1.Hour(),
			time1.Minute(),
			time1.Second())

		_, err = file.WriteString("Username: " + username2 + "\nPassword: " + password2 + "\nTime: " + time2)
		if err != nil {
			log.Fatal(err)
		}

		if status {
			break
		}

	}

}
