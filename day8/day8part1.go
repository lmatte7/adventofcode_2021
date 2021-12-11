package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	_, outputValues := parseInput("input.txt")

	uniqueDigits := 0
	for _, output := range outputValues {

		values := strings.Split(output, " ")

		for _, value := range values {
			fmt.Printf("Value: %s\n", value)

			if len(value) == 0 {
				continue
			}

			if len(value) == 2 || len(value) == 4 || len(value) == 3 || len(value) == 7 {
				uniqueDigits++
			}

		}
		fmt.Printf("\n")
	}

	fmt.Printf("Unique Digits: %d", uniqueDigits)
}

func parseInput(filename string) (signalPatterns []string, outputValues []string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		splitInput := strings.Split(scanner.Text(), "|")

		signalPatterns = append(outputValues, splitInput[0])
		outputValues = append(outputValues, splitInput[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return signalPatterns, outputValues
}
