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

	input := parseInput("input.txt")

	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, command := range input {

		commandComponents := strings.Split(command, " ")
		commandValue, err := strconv.Atoi(commandComponents[1])
		if err != nil {
			panic(err)
		}

		switch commandComponents[0] {

		case "forward":
			horizontalPosition += commandValue
			depth += aim * commandValue
		case "down":
			aim += commandValue
		case "up":
			aim -= commandValue
		default:
			fmt.Printf("UNKNOWN COMMAND: %s\n", commandComponents[0])
		}

	}

	finalPosition := horizontalPosition * depth

	fmt.Println(finalPosition)

}

func parseInput(filename string) []string {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	values := make([]string, 0)

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return values
}
