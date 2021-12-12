package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	steps := 900
	powers := parseInput("input.txt")
	flashes := 0

	for k := steps; k > 0; k-- {

		flashedOctopi := make([][]int, len(powers))
		for i := 0; i < len(powers); i++ {
			flashedOctopi[i] = make([]int, len(powers[i]))
		}
		for i := 0; i < len(powers); i++ {

			for j := 0; j < len(powers[i]); j++ {
				powers = checkPower(i, j, powers, &flashes, &flashedOctopi)
			}
		}

		allMatch := true
		for k := 0; k < len(powers); k++ {
			if !allMatch {
				break
			}
			for l := 0; l < len(powers[k]); l++ {
				if powers[k][l] != powers[0][0] {
					allMatch = false
				}
			}
		}

		if allMatch {
			fmt.Printf("Sync at: %d\n", (steps - k + 1))
		}

	}

	fmt.Printf("Total Flases: %d", flashes)
}

func checkPower(i int, j int, powers [][]int, flashes *int, flashedOctopi *[][]int) [][]int {
	fOPointer := *flashedOctopi

	if fOPointer[i][j] != 1 {

		powers[i][j]++
		if powers[i][j] > 9 {

			// for k := 0; k < len(powers); k++ {

			// 	for l := 0; l < len(powers[k]); l++ {
			// 		if k == i && l == j {
			// 			fmt.Printf("%s", "X")
			// 		} else {
			// 			fmt.Printf("%d", powers[k][l])
			// 		}
			// 	}
			// 	fmt.Println()
			// }
			// fmt.Println()
			// fmt.Println()

			powers[i][j] = 0
			fOPointer[i][j] = 1
			*flashes++

			if j-1 >= 0 {
				powers = checkPower(i, j-1, powers, flashes, flashedOctopi)
			}
			if i-1 >= 0 {
				powers = checkPower(i-1, j, powers, flashes, flashedOctopi)
			}
			if j+1 < len(powers[i]) {
				powers = checkPower(i, j+1, powers, flashes, flashedOctopi)
			}
			if i+1 < len(powers) {
				powers = checkPower(i+1, j, powers, flashes, flashedOctopi)
			}
			if i+1 < len(powers) && j+1 < len(powers[i]) {
				powers = checkPower(i+1, j+1, powers, flashes, flashedOctopi)
			}
			if i-1 >= 0 && j-1 >= 0 {
				powers = checkPower(i-1, j-1, powers, flashes, flashedOctopi)
			}
			if i+1 < len(powers) && j-1 >= 0 {
				powers = checkPower(i+1, j-1, powers, flashes, flashedOctopi)
			}
			if i-1 >= 0 && j+1 < len(powers[i]) {
				powers = checkPower(i-1, j+1, powers, flashes, flashedOctopi)
			}

		}
	}
	return powers
}

func parseInput(filename string) (powers [][]int) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	column := 0
	for scanner.Scan() {

		row := scanner.Text()
		powers = append(powers, make([]int, 0))

		for _, value := range strings.Split(row, "") {

			number, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			powers[column] = append(powers[column], number)
		}

		column++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return powers
}
