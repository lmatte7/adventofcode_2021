package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	input := parseInput("input.txt")

	depthIncreases := 0

	for key, measurement := range input {

		if key > 0 && measurement > input[key-1] {
			depthIncreases++
		}
	}

	fmt.Println(depthIncreases)
}

func parseInput(filename string) []int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	values := make([]int, 0)

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		intValue, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		values = append(values, intValue)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values
}
