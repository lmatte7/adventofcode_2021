package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	signals := parseInput("input.txt")

	for _, signal := range signals {

		splitInput := strings.Split(signal, "|")

		patterns := strings.Split(splitInput[0], " ")

		sort.Slice(patterns, func(i, j int) bool {
			return len(patterns[i]) < len(patterns[j])
		})

		letters := findLetters(patterns, false)

		for _, output := range patterns {

			if len(output) == 6 && strings.Contains(output, letters[4]) {
				if strings.Contains(output, letters[0]) && strings.Contains(output, letters[1]) && strings.Contains(output, letters[3]) && strings.Contains(output, letters[4]) && strings.Contains(output, letters[2]) && strings.Contains(output, letters[6]) {
					letters = findLetters(patterns, true)
				}
			}
		}

		outputVales := strings.Split(splitInput[1], " ")
		outputNumber := ""
		// fmt.Printf("LEtters: %v\n", letters)
		// fmt.Printf("Output Values: %v\n", outputVales)

		for _, value := range outputVales {
			if len(value) == 0 {
				continue
			}
			// fmt.Printf("Value: %s\n", value)

			if len(value) == 7 && strings.Contains(value, letters[1]) && strings.Contains(value, letters[5]) && strings.Contains(value, letters[0]) && strings.Contains(value, letters[3]) && strings.Contains(value, letters[6]) && strings.Contains(value, letters[2]) && strings.Contains(value, letters[4]) {
				outputNumber += "8"
			} else if len(value) == 6 && strings.Contains(value, letters[1]) && strings.Contains(value, letters[5]) && strings.Contains(value, letters[0]) && strings.Contains(value, letters[3]) && strings.Contains(value, letters[6]) && strings.Contains(value, letters[2]) {
				outputNumber += "9"
			} else if len(value) == 6 && strings.Contains(value, letters[0]) && strings.Contains(value, letters[1]) && strings.Contains(value, letters[3]) && strings.Contains(value, letters[4]) && strings.Contains(value, letters[5]) && strings.Contains(value, letters[6]) {
				outputNumber += "6"
			} else if len(value) == 5 && strings.Contains(value, letters[0]) && strings.Contains(value, letters[3]) && strings.Contains(value, letters[4]) && strings.Contains(value, letters[5]) && strings.Contains(value, letters[6]) {
				outputNumber += "2"
			} else if len(value) == 5 && strings.Contains(value, letters[2]) && strings.Contains(value, letters[5]) && strings.Contains(value, letters[0]) && strings.Contains(value, letters[3]) && strings.Contains(value, letters[6]) {
				outputNumber += "3"
			} else if len(value) == 5 && strings.Contains(value, letters[1]) && strings.Contains(value, letters[5]) && strings.Contains(value, letters[0]) && strings.Contains(value, letters[3]) && strings.Contains(value, letters[6]) {
				outputNumber += "5"
			} else if len(value) == 4 && strings.Contains(value, letters[2]) && strings.Contains(value, letters[5]) && strings.Contains(value, letters[3]) && strings.Contains(value, letters[1]) {
				outputNumber += "4"
			} else if len(value) == 3 && strings.Contains(value, letters[2]) && strings.Contains(value, letters[5]) && strings.Contains(value, letters[0]) {
				outputNumber += "7"
			} else if len(value) == 2 && strings.Contains(value, letters[2]) && strings.Contains(value, letters[5]) {
				outputNumber += "1"
			}

			// fmt.Printf("Current Number: %s\n", outputNumber)
		}

		fmt.Println(outputNumber)

	}

}

func parseInput(filename string) (signals []string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		signals = append(signals, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return signals
}

func findLetters(patterns []string, swap bool) (letters []string) {

	letters = []string{" ", " ", " ", " ", " ", " ", " "}
	for _, output := range patterns {

		if len(output) == 0 {
			continue
		}

		if len(output) == 2 {
			if swap {
				letters[2] = string(output[0])
				letters[5] = string(output[1])
			} else {
				letters[5] = string(output[0])
				letters[2] = string(output[1])
			}
		}

		if len(output) == 3 {
			for _, letter := range strings.Split(output, "") {
				if letter != letters[2] && letter != letters[5] {
					letters[0] = letter
				}
			}
		}

		if len(output) == 4 {

			patternFour := output
			for _, subOutput := range patterns {
				if len(subOutput) == 5 {

					if strings.Contains(subOutput, letters[0]) && strings.Contains(subOutput, letters[2]) && strings.Contains(subOutput, letters[5]) {
						for _, letter := range strings.Split(subOutput, "") {
							if strings.Contains(patternFour, letter) && letter != letters[2] && letter != letters[5] {
								letters[3] = letter
							}
							if !strings.Contains(patternFour, letter) && letter != letters[0] {
								letters[6] = letter
							}
							for _, patternFourLetter := range strings.Split(patternFour, "") {
								if !strings.Contains(subOutput, patternFourLetter) {
									letters[1] = patternFourLetter
								}
							}
						}
					}

				}
			}
		}

		if len(output) == 7 {
			for _, letter := range strings.Split(output, "") {
				if letter != letters[0] && letter != letters[1] && letter != letters[2] && letter != letters[3] && letter != letters[5] && letter != letters[6] {
					letters[4] = letter
				}
			}
		}
	}

	return letters
}
