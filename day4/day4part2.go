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
			columnFarthestHold = make([]int, 5)
		}

	}

	fmt.Printf("Farther Row Number Positions: %v\n", farthestRowNumbers)
	fmt.Printf("Farther Column Number Positions: %v\n", farthestColumnNumbers)

	lowestRowNumber := len(positions)
	lowestColumnNumber := len(positions)

	lowestRowNumberKey := 0
	lowestColumnNumberKey := 0

	lowestRowBoardNumberPositions := make([]int, (len(boards) / 5))
	lowestColumnBoardNumberPositions := make([]int, (len(boards) / 5))

	for i := 0; i < len(lowestRowBoardNumberPositions); i++ {
		lowestColumnBoardNumberPositions[i] = 99
		lowestRowBoardNumberPositions[i] = 99
	}

	for i := 0; i < len(boards); i += 5 {

		for j := 0; j < 5; j++ {
			if farthestRowNumbers[i+j] < lowestRowBoardNumberPositions[i/5] {
				lowestRowBoardNumberPositions[i/5] = farthestRowNumbers[i+j]
			}
			if farthestColumnNumbers[i+j] < lowestColumnBoardNumberPositions[i/5] {
				lowestColumnBoardNumberPositions[i/5] = farthestColumnNumbers[i+j]
			}
		}
	}

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

	lowestBoardNumberPositions := make([]int, (len(boards) / 5))

	for key := range lowestRowBoardNumberPositions {
		if lowestColumnBoardNumberPositions[key] < lowestRowBoardNumberPositions[key] {
			lowestBoardNumberPositions[key] = lowestColumnBoardNumberPositions[key]
		} else {
			lowestBoardNumberPositions[key] = lowestRowBoardNumberPositions[key]
		}
	}

	fmt.Printf("Lowest Row Board Numbers: %v\n", lowestRowBoardNumberPositions)
	fmt.Printf("Lowest Column Board Numbers: %v\n", lowestColumnBoardNumberPositions)
	fmt.Printf("Lowest Overall Board Numbers: %v\n", lowestBoardNumberPositions)

	fmt.Printf("Lowest Row: %d, Lowest Column: %d\n", lowestRowNumber, lowestColumnNumber)
	fmt.Printf("Lowest Row Key: %d, Lowest Column Key: %d\n", lowestRowNumberKey, lowestColumnNumberKey)

	newLowestBoard := 0
	newLowestNumber := 0
	for key, number := range lowestBoardNumberPositions {
		if number > newLowestNumber {
			newLowestBoard = key
			newLowestNumber = number
		}
	}

	fmt.Println(newLowestBoard)
	fmt.Println(newLowestNumber)

	uncalledNumberSum := 0
	for i := (newLowestBoard * 5); i < ((newLowestBoard + 1) * 5); i++ {

		for _, number := range boards[i] {
			if positions[number] > newLowestNumber {
				uncalledNumberSum += number
			}
		}
	}

	fmt.Printf("Sum of uncalled Numbers: %d\n", uncalledNumberSum)
	fmt.Printf("Last Number: %d\n", intNumbers[newLowestNumber])
	fmt.Printf("Result: %d\n", intNumbers[newLowestNumber]*uncalledNumberSum)
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
