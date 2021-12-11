package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// [2 1 9 9 9 4 3 2 1 0]
// [3 9 8 7 8 9 4 9 2 1]
// [9 8 5 6 7 8 9 8 9 2]
// [8 7 6 7 8 9 6 7 8 9]
// [9 8 9 9 9 6 5 6 7 8]

func main() {

	heights := parseInput("input.txt")

	riskTotal := 0
	for i := 0; i < len(heights); i++ {

		for j := 0; j < len(heights[i]); j++ {

			isLower := true
			// fmt.Printf("Current Number: %d ", heights[i][j])
			// if i-1 >= 0 {
			// 	fmt.Printf("Above: %d", heights[i-1][j])
			// }
			// fmt.Println()
			if j-1 >= 0 && heights[i][j] >= heights[i][j-1] {
				isLower = false
			}
			if i-1 >= 0 && heights[i][j] >= heights[i-1][j] {
				isLower = false
			}
			if j+1 < len(heights[i]) && heights[i][j] >= heights[i][j+1] {
				isLower = false

			}
			if i+1 < len(heights) && heights[i][j] >= heights[i+1][j] {
				isLower = false

			}
			if isLower {
				fmt.Printf("Number %d at [%d,%d] is lower\n", heights[i][j], i, j)
				riskTotal += 1 + heights[i][j]
			}
		}

	}
	fmt.Printf("Risk Total: %d\n", riskTotal)
}
func parseInput(filename string) (heights [][]int) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	column := 0
	for scanner.Scan() {

		row := scanner.Text()
		heights = append(heights, make([]int, 0))

		for _, value := range strings.Split(row, "") {

			number, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			heights[column] = append(heights[column], number)
		}

		column++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return heights
}
