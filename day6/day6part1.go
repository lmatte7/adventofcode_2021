package main

import "fmt"

type lanterFish struct {
	Timer int
	// Children int
}

func main() {
	fmt.Println(Solve(GetInput(), 18))
}

func Solve(days [9]int, n int) int {
	for i := 0; i < n; i++ {
		days[0], days[1], days[2], days[3], days[4], days[5], days[6], days[7], days[8] = days[1], days[2], days[3], days[4], days[5], days[6], days[0]+days[7], days[8], days[0]
		fmt.Println(days)
	}
	return days[0] + days[1] + days[2] + days[3] + days[4] + days[5] + days[6] + days[7] + days[8]
}

func GetInput() [9]int {
	var days [9]int
	for _, item := range []uint8{3, 4, 3, 1, 2} {
		days[item]++
	}
	return days
}
