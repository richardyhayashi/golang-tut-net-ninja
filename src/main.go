package main

import (
	"fmt"
)

// Entry point of app.
func main() {
	// Simple loop
	/* x := 0
	for x < 5 {
		fmt.Println("Value of x is:", x)
		x++
	} */

	// For loop
	/* for i := 0; i < 5; i++ {
		fmt.Println("Value of x is:", i)
	} */

	names := []string{"Mario", "Luigi", "Yoshi", "Peach"}

	/* for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	} */

	/* for index, value := range names {
		fmt.Printf("The value at index %v is %v\n", index, value)
	} */

	for _, value := range names {
		fmt.Printf("The value is %v\n", value)
		value = "new string"
	}

	fmt.Println(names)
}
