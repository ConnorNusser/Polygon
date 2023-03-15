package main

import (
	"fmt"
	"sort"
)

func main() {
	/* var nameThree string
	fmt.Println("HelloTest")

	var nameOne string = "mario"
	var nameTwo = "luigi"
	fmt.Println(nameOne)
	fmt.Println(nameTwo)

	nameFour := "yoshi" */

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	var ages = []int{20, 25, 30}
	fmt.Println(ages[0])
	ages = append(ages, 85)

	rangeOne := ages[1:3]
	fmt.Println(rangeOne)

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)
}
