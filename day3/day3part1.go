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

	bitCount := make([]int, len(input[0]))

	for _, bNum := range input {

		for i := 0; i < len(bNum); i++ {
			current, err := strconv.Atoi(string(bNum[i]))
			if err != nil {
				panic(err)
			}
			if current == 1 {
				bitCount[i] += current
			} else {
				bitCount[i] -= 1
			}
		}
	}

	mcb := ""
	lcb := ""

	for i := 0; i < len(bitCount); i++ {
		if bitCount[i] > 0 {
			mcb += "1"
			lcb += "0"
		} else {
			mcb += "0"
			lcb += "1"
		}
	}

	mcbBinary, err := strconv.ParseInt(mcb, 2, 64)
	if err != nil {
		panic(err)
	}
	lcbBinary, err := strconv.ParseInt(lcb, 2, 64)
	if err != nil {
		panic(err)
	}

	result := int(mcbBinary) * int(lcbBinary)

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
