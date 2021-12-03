package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	baseInput := parseInput("input.txt")

	bitCount := make([]int, len(baseInput[0]))

	input := baseInput
	for i := 0; i < len(input[0]); i++ {

		remainingHighMeasurements := make([]string, 0)
		remainingLowMeasurements := make([]string, 0)

		for j := 0; j < len(input); j++ {

			current, err := strconv.Atoi(string(input[j][i]))
			if err != nil {
				panic(err)
			}
			if current == 1 {
				bitCount[i] += current
				remainingHighMeasurements = append(remainingHighMeasurements, input[j])
			} else {
				bitCount[i] -= 1
				remainingLowMeasurements = append(remainingLowMeasurements, input[j])
			}
		}

		if bitCount[i] >= 0 {
			input = remainingHighMeasurements
		} else {
			input = remainingLowMeasurements
		}

		if len(input) == 1 {
			break
		}
	}

	oxygenRating := input[0]

	input = baseInput
	bitCount = make([]int, len(baseInput[0]))
	for i := 0; i < len(input[0]); i++ {

		remainingHighMeasurements := make([]string, 0)
		remainingLowMeasurements := make([]string, 0)

		for j := 0; j < len(input); j++ {

			current, err := strconv.Atoi(string(input[j][i]))
			if err != nil {
				panic(err)
			}
			if current == 1 {
				bitCount[i] += current
				remainingLowMeasurements = append(remainingLowMeasurements, input[j])
			} else {
				bitCount[i] -= 1
				remainingHighMeasurements = append(remainingHighMeasurements, input[j])
			}
		}

		if bitCount[i] >= 0 {
			input = remainingHighMeasurements
		} else {
			input = remainingLowMeasurements
		}

		if len(input) == 1 {
			break
		}
	}

	co2Rating := input[0]

	oxygenBinary, err := strconv.ParseInt(oxygenRating, 2, 64)
	if err != nil {
		panic(err)
	}
	co2Binary, err := strconv.ParseInt(co2Rating, 2, 64)
	if err != nil {
		panic(err)
	}

	result := int(oxygenBinary) * int(co2Binary)

	fmt.Printf("Result: %d\n", result)

}

func parseInput(filename string) []string {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	values := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values
}
