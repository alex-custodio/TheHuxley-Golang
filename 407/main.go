package main

import "fmt"

func main(){
	var i int
	var j int
	var max int
	fmt.Scanln(&i, &j)
	max = MaxCycleSize(i,j)
	fmt.Println(i, j, max)
}

func MaxCycleSize(initialNumber, finalNumber int) int {
	var calc int
	var max int = 0
	for initialNumber <= finalNumber {
		calc = calc3n1(initialNumber)
		cycleSizeCount := 0
		for calc != 1 {
			calc = calc3n1(calc)
			cycleSizeCount++
		}
		if cycleSizeCount > max {
			max = cycleSizeCount
		}
		initialNumber += 1
	}
	// Final number and initial number added to max
	max += 2
	return max
}

func calc3n1(number int) int {
	var calc int
	if number % 2 == 0 {
		calc = number / 2
	} else {
		calc = number * 3 + 1
	}
	return calc
}
