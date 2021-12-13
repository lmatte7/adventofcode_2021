package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	points, folds := parseInput("input.txt")

	// fmt.Println(points)
	// fmt.Println(folds)
	// fmt.Println("First Half")
	// for _, point := range points {
	// 	fmt.Println(point)
	// }

	maxX := 0
	maxY := 0
	finalPoints := points
	for _, fold := range folds {
		sections := strings.Split(fold, " ")

		coordinate := strings.Split(sections[2], "=")
		foldNumber, err := strconv.Atoi(string(coordinate[1]))
		if err != nil {
			panic(err)
		}

		fmt.Printf("Fold Number: %d\n", foldNumber)

		firstHalf := make([][]int, 0)
		secondHalf := make([][]int, 0)

		for _, point := range finalPoints {

			if string(coordinate[0]) == "y" {
				maxY = foldNumber
				if point[1] < foldNumber {
					firstHalf = append(firstHalf, point)
				} else {
					newY := math.Abs(float64((point[1] - (foldNumber)) - foldNumber))
					point = []int{point[0], int(newY)}
					secondHalf = append(secondHalf, point)
				}
			} else {
				maxX = foldNumber
				if point[0] < foldNumber {
					firstHalf = append(firstHalf, point)
				} else {
					newX := math.Abs(float64((point[0] - (foldNumber)) - foldNumber))
					point = []int{int(newX), point[1]}
					secondHalf = append(secondHalf, point)
				}
			}

		}
		finalPoints = firstHalf

		for _, secondPoint := range secondHalf {
			inOtherArray := false

			for _, firstPoint := range finalPoints {
				if firstPoint[0] == secondPoint[0] && firstPoint[1] == secondPoint[1] {
					inOtherArray = true
				}

			}

			if !inOtherArray {
				finalPoints = append(finalPoints, secondPoint)
			}
		}

		// fmt.Println("First Half")
		// for _, point := range firstHalf {
		// 	fmt.Println(point)
		// }
		// fmt.Println("Second Half")
		// for _, point := range secondHalf {
		// 	fmt.Println(point)
		// }
		// fmt.Println("Final Points")
		// for _, point := range finalPoints {
		// 	fmt.Println(point)
		// }

	}

	for _, point := range finalPoints {
		if point[0] > maxX {
			maxX = point[0]
		}
		if point[1] > maxY {
			maxY = point[1]
		}
	}

	pointDisplay := make([][]string, maxY)

	fmt.Printf("Max X: %d, Max Y: %d\n", maxX, maxY)
	for i := 0; i < maxY; i++ {
		pointDisplay[i] = make([]string, maxX)
	}

	for _, point := range finalPoints {
		pointDisplay[point[1]][point[0]] = "X"
	}

	for _, row := range pointDisplay {

		for _, column := range row {
			if column == "" {
				fmt.Printf("%s", " ")
			} else {
				fmt.Printf("%s", column)
			}
		}
		fmt.Println()
	}
}

func parseInput(filename string) (points [][]int, folds []string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	column := 0
	for scanner.Scan() {

		line := scanner.Text()
		foldLine, err := regexp.MatchString("^fold*", line)
		if err != nil {
			panic(err)
		}
		if foldLine {
			folds = append(folds, line)
		} else {
			if len(line) < 2 {
				continue
			}
			points = append(points, make([]int, 0))

			for _, value := range strings.Split(line, ",") {

				number, err := strconv.Atoi(value)
				if err != nil {
					panic(err)
				}

				points[column] = append(points[column], number)
			}

			column++
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return points, folds
}
