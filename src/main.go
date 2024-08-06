package main

import (
	"fmt"
	"sort"
)

// Entry point of app.
func main() {
	//greeting := "Hello there my friends!"

	//fmt.Println(strings.Contains(greeting, "Hello"))
	//fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	//fmt.Println(strings.ToUpper(greeting))
	//fmt.Println(strings.Index(greeting, "th"))
	//fmt.Println(strings.Split(greeting, " "))

	// The original value is unchanged.
	//fmt.Println("Original string value =", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 90)
	fmt.Println(index)

	names := []string{"Yoshi", "Mario", "Peach", "Bowser", "Luigi"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Bowser"))
}
