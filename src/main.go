package main

import (
	"fmt"
)

// Entry point of app.
func main() {
	age := 25

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less then 30")	
	} else if age < 40 {
		fmt.Println("Age is less then 40")	
	} else {
		fmt.Println("Age is NOT less then 45")	
	}

	names := []string{"Mario", "Luigi", "Yoshi", "Peach", "Bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		} 

		if index > 2 {
			fmt.Println("Break at pos", index)
			break
		}

		fmt.Printf("The value at pos %v is %v\n", index, value)
	}
}
