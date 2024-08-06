package main

import "fmt"

// Entry point of app.
func main() {
	age := 35
	name := "John"

	// Pint
	fmt.Print("Hello")
	fmt.Print(", World!\n")
	fmt.Print("New line.\n")

	// Println
	fmt.Println("Hello World")
	fmt.Println("Goodbye World!")
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v\n", age, name)
	fmt.Printf("My age is %q and my name is %q\n", age, name)
	fmt.Printf("Age is of type %T\n", age)
	fmt.Printf("Your scored %f points!\n", 255.55)
	fmt.Printf("Your scored %0.1f points!\n", 255.55)

	// Sprintf (Save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v", age, name)
	fmt.Println("The saved string is:", str)
}
