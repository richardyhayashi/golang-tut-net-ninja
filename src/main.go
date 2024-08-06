package main

import "fmt"

// Entry point of app.
func main() {
	//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}
	
	names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	names[1] = "Luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Koopa")

	fmt.Println(rangeOne)
}
