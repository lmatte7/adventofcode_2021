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

	stringNumbers, boards := parseInput("input.txt")

	fmt.Printf("%v\n", stringNumbers)
	positions := make(map[int]int)

	numbers := strings.Split(stringNumbers, ",")
	intNumbers := make([]int, len(numbers))

	for key, number := range numbers {
		currentNum, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		positions[currentNum] = key
		intNumbers[key] = currentNum
	}

	fmt.Printf("%v\n", positions)
	// rowsAndColumns := make([]int, 0)

	farthestRowNumbers := make([]int, 0)
	farthestColumnNumbers := make([]int, 0)
	columnFarthestHold := make([]int, 5)

	columnCount := 0

	for _, board := range boards {
		fmt.Printf("%v\n", board)

		highestRowNumber := 0
		for key, number := range board {
			if positions[number] > highestRowNumber {
				highestRowNumber = positions[number]
			}

			if positions[number] > columnFarthestHold[key] {
				columnFarthestHold[key] = positions[number]
			}
		}

		farthestRowNumbers = append(farthestRowNumbers, highestRowNumber)

		if columnCount < 4 {
			columnCount++
		} else {
			columnCount = 0
			for _, number := range columnFarthestHold {
				farthestColumnNumbers = append(farthestColumnNumbers, number)
			}
		}

	}

	fmt.Printf("Farther Row Number Positions: %v\n", farthestRowNumbers)
	fmt.Printf("Farther Column Number Positions: %v\n", farthestColumnNumbers)

	lowestRowNumber := len(positions)
	lowestColumnNumber := len(positions)

	lowestRowNumberKey := 0
	lowestColumnNumberKey := 0

	for key, number := range farthestRowNumbers {
		if number < lowestRowNumber {
			lowestRowNumber = number
			lowestRowNumberKey = key
		}
	}

	for key, number := range farthestColumnNumbers {
		if number < lowestColumnNumber {
			lowestColumnNumber = number
			lowestColumnNumberKey = key
		}
	}

	winningBoard := 0
	lowestNumber := 0

	fmt.Printf("Lowest Row: %d, Lowest Column: %d\n", lowestRowNumber, lowestColumnNumber)
	fmt.Printf("Lowest Row Key: %d, Lowest Column Key: %d\n", lowestRowNumberKey, lowestColumnNumberKey)
	if lowestRowNumber < lowestColumnNumber {
		winningBoard = lowestRowNumberKey / 5
		lowestNumber = lowestRowNumber
	} else {
		winningBoard = lowestColumnNumberKey / 5
		lowestNumber = lowestColumnNumber
	}

	fmt.Println(winningBoard)

	uncalledNumberSum := 0
	for i := (winningBoard * 5); i < ((winningBoard + 1) * 5); i++ {

		for _, number := range boards[i] {
			if positions[number] > lowestNumber {
				uncalledNumberSum += number
			}
		}
	}

	fmt.Printf("Sum of uncalled Numbers: %d\n", uncalledNumberSum)
	fmt.Printf("Last Number: %d\n", intNumbers[lowestNumber])
	fmt.Printf("Result: %d\n", intNumbers[lowestNumber]*uncalledNumberSum)
}

func parseInput(filename string) (numbers string, boards [][]int) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	boards = make([][]int, 0)

	boardCounter := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if len(scanner.Text()) > 20 {
			numbers = scanner.Text()
		} else if len(scanner.Text()) > 4 {

			numbers := strings.Split(scanner.Text(), " ")

			intNumberArray := make([]int, 0)
			for _, number := range numbers {
				if len(number) == 0 {
					continue
				}
				currentNum, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}

				intNumberArray = append(intNumberArray, currentNum)
			}

			boards = append(boards, intNumberArray)
		} else {
			boardCounter++
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers, boards
}
