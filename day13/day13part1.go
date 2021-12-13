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
	part1 := true

	// fmt.Println(points)
	// fmt.Println(folds)
	// fmt.Println("First Half")
	// for _, point := range points {
	// 	fmt.Println(point)
	// }

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

		for _, point := range points {

			if string(coordinate[0]) == "y" {
				if point[1] < foldNumber {
					firstHalf = append(firstHalf, point)
				} else {
					newY := math.Abs(float64((point[1] - (foldNumber)) - foldNumber))
					point = []int{point[0], int(newY)}
					secondHalf = append(secondHalf, point)
				}
			} else {
				if point[0] < foldNumber {
					firstHalf = append(firstHalf, point)
				} else {
					newX := math.Abs(float64((point[0] - (foldNumber)) - foldNumber))
					point = []int{int(newX), point[1]}
					secondHalf = append(secondHalf, point)
				}
			}

		}
		finalPoints := firstHalf

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
		fmt.Printf("Final Points: %d\n", len(finalPoints))
		// fmt.Println("Final Points")
		// for _, point := range finalPoints {
		// 	fmt.Println(point)
		// }

		if part1 {
			break
		}
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
